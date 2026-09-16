"""Check locale coverage, shared snippets, and optionally all rendered local links."""
import argparse
from html.parser import HTMLParser
from pathlib import Path
import posixpath
import re
from urllib.parse import unquote, urlsplit

ROOT = Path(__file__).resolve().parents[1]
DOCS = ROOT / 'docs'


class Page(HTMLParser):
    def __init__(self, path):
        super().__init__()
        self.ids = set()
        self.links = []
        self.lang = None
        self.feed(path.read_text(encoding='utf-8'))

    def handle_starttag(self, tag, attrs):
        attrs = dict(attrs)
        if 'id' in attrs:
            self.ids.add(attrs['id'])
        if tag == 'a' and 'href' in attrs:
            self.links.append(attrs['href'])
        if tag == 'html':
            self.lang = attrs.get('lang')


def check_sources():
    english = {p.relative_to(DOCS).as_posix(): p for p in DOCS.rglob('*.md')
               if '.vitepress' not in p.parts and 'zh-cn' not in p.relative_to(DOCS).parts}
    chinese = {p.relative_to(DOCS / 'zh-cn').as_posix(): p for p in (DOCS / 'zh-cn').rglob('*.md')}
    assert english.keys() == chinese.keys(), 'English and Chinese page coverage differs'
    for route, en in english.items():
        text = en.read_text(encoding='utf-8')
        assert not re.search(r'[\u3400-\u9fff]', text), f'Untranslated English content: {en}'
        examples = []
        for page in (en, chinese[route]):
            paths = [(page.parent / path).resolve() for path in
                     re.findall(r'^<<<\s+(\S+)', page.read_text(encoding='utf-8'), re.M)]
            for path in paths:
                assert path.is_file(), f'Missing snippet: {page} -> {path}'
                assert (DOCS / 'examples').resolve() in path.parents, f'Unshared snippet: {path}'
            examples.append(paths)
        assert examples[0] == examples[1], f'Locales include different examples: {route}'
    print(f'Locale coverage: {len(english)} pages per language; all snippets shared.')


def check_dist():
    dist = DOCS / '.vitepress/dist'
    pages = {p.relative_to(dist).as_posix(): Page(p) for p in dist.rglob('*.html')}
    assert pages, 'Build the site before using --dist'
    errors = []
    links = 0
    for name, page in pages.items():
        assert page.lang == ('zh-CN' if name.startswith('zh-cn/') else 'en'), f'Wrong lang: {name}'
        for link in page.links:
            url = urlsplit(link)
            if url.scheme or url.netloc:
                continue
            path = unquote(url.path)
            target = posixpath.normpath(path.lstrip('/') if path.startswith('/') else
                                       posixpath.join(posixpath.dirname(name), path)) if path else name
            if target == '.':
                target = 'index.html'
            key = next((p for p in (target, target + '.html', target.rstrip('/') + '/index.html')
                        if p in pages), None)
            if key is None:
                if not (dist / target).is_file():
                    errors.append(f'{name}: missing target {link}')
            elif url.fragment and unquote(url.fragment) not in pages[key].ids:
                errors.append(f'{name}: missing anchor {link}')
            links += 1
    assert not errors, '\n'.join(errors)
    assert (dist / 'llms.txt').read_bytes() == (DOCS / 'public/llms.txt').read_bytes()
    print(f'Rendered locales and {links} local links verified.')


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--dist', action='store_true')
    args = parser.parse_args()
    check_sources()
    if args.dist:
        check_dist()
