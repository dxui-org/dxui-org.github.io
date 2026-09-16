"""Check API drift and compile every published complete Go example without opening windows."""
import json
import os
import re
from pathlib import Path
import shutil
import subprocess
import sys
import tempfile

root = Path(__file__).resolve().parents[1]
source = Path(sys.argv[1]).resolve()
env = {**os.environ, 'CGO_ENABLED': '0', 'GOWORK': 'off'}
subprocess.run(['go', 'run', './scripts/format-examples.go', '-check', 'docs'], cwd=root, env=env, check=True)
result = subprocess.run(['go', 'run', './scripts/api-scan.go', str(source)], cwd=root,
                        env=env, check=True, capture_output=True)
actual = json.loads(result.stdout)
expected = json.loads((root/'scripts/api-snapshot.json').read_text(encoding='utf-8-sig'))
if actual != expected:
    raise SystemExit('API snapshot differs. Regenerate it and review the Chinese explanations.')
print(f'API snapshot matches: {len(actual)} declarations.', flush=True)
reference = '\n'.join(p.read_text(encoding='utf-8') for p in (root/'docs/api').glob('*.md'))
for entry in actual:
    if entry['Package'] == '.' and entry['Name'] not in reference:
        raise SystemExit(f'Missing reference entry: {entry["Name"]}')
    if entry['Package'] == 'icon' and f'icon.{entry["Name"]}' not in reference:
        raise SystemExit(f'Missing icon entry: {entry["Name"]}')
known = {entry['Name'].split('.')[0] for entry in actual if entry['Package'] == '.'}
for file in (root/'docs').rglob('*.md'):
    if '.vitepress' in file.parts: continue
    if re.search(r'\bsdl\w*', file.read_text(encoding='utf-8'), re.IGNORECASE):
        raise SystemExit(f'Unexpected backend implementation details in {file}')
    for symbol in re.findall(r'\bdxui\.([A-Z]\w*)', file.read_text(encoding='utf-8')):
        if symbol not in known:
            raise SystemExit(f'Unknown public name dxui.{symbol} in {file}')
examples = sorted((root/'docs/examples').glob('*/main.go'))
if not examples: raise SystemExit('No examples found')
for file in examples:
    formatted = subprocess.run(['gofmt', '-l', str(file)], check=True, capture_output=True).stdout.strip()
    if formatted: raise SystemExit(f'Run gofmt on {file}')
with tempfile.TemporaryDirectory(prefix='.docs-check-', dir=root) as tmp:
    tmp = Path(tmp)
    if root.resolve() not in tmp.resolve().parents:
        raise SystemExit('Temporary build path is outside the documentation workspace')
    mod = (source/'go.mod').read_text().replace('module github.com/dxui-org/dxui', 'module example.com/dxui-docs', 1)
    mod += f'\nrequire github.com/dxui-org/dxui v0.0.0\nreplace github.com/dxui-org/dxui => "{source.as_posix()}"\n'
    (tmp/'go.mod').write_text(mod)
    shutil.copyfile(source/'go.sum', tmp/'go.sum')
    for file in examples:
        target = tmp/file.parent.name
        target.mkdir()
        shutil.copyfile(file, target/'main.go')
    subprocess.run(['go', 'build', '-mod=mod', './...'], cwd=tmp, env=env, check=True)
print(f'Compiled {len(examples)} complete examples with CGO_ENABLED=0.', flush=True)
