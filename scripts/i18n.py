"""Shared localization helpers for generated reference pages.

Handwritten pages live directly in docs (English) and docs/zh-cn (Chinese).
Generated Chinese text is translated using reviewed entries in i18n-en.json.
Unknown translations fail generation rather than leaking Chinese into English pages.
"""
import json
from pathlib import Path
import posixpath
import re

DATA = json.loads(Path(__file__).with_name('i18n-en.json').read_text(encoding='utf-8'))
HAN = re.compile(r'[\u3400-\u9fff]')


def english(text):
    result = []
    for original in text.splitlines():
        line = original
        if HAN.search(line):
            for source, target in DATA['phrases'].items():
                line = line.replace(source, target)
            if HAN.search(line):
                if line not in DATA['lines']:
                    raise ValueError(f'Missing English translation: {original}')
                line = DATA['lines'][line]
            line = line.replace('：', ': ').replace('；', '; ').replace('。', '.')
        result.append(line)
    return '\n'.join(result) + '\n'


def chinese(path, text):
    # Local documentation links stay in Chinese. Shared assets and llms.txt stay at root.
    text = re.sub(r'(?<=[(" ])/(guide|components|api|ai|examples)(?=[/#"\s)]|$)',
                  r'/zh-cn/\1', text)
    def snippet(match):
        source = posixpath.normpath(posixpath.join(posixpath.dirname(path), match[1]))
        relative = posixpath.relpath(source, posixpath.join('zh-cn', posixpath.dirname(path)))
        return '<<< ' + relative
    return re.sub(r'^<<<\s+(\S+)', snippet, text, flags=re.M)
