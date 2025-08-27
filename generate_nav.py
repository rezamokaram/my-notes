import os
import sys

def build_nav(dir_path, indent=2):
    nav_lines = []
    items = sorted(os.listdir(dir_path))
    
    for item in items:
        full_path = os.path.join(dir_path, item)
        
        if os.path.isdir(full_path):
            # Directory: recurse
            nav_lines.append(' ' * indent + f"- {item}:")
            nav_lines.extend(build_nav(full_path, indent + 2))
        
        elif item.endswith('.md'):
            # Markdown file: list without extension in nav
            # You can use full relative path from docs_dir here
            rel_path = os.path.relpath(full_path, start='notes')
            nav_lines.append(' ' * indent + f"- {item}: {rel_path.replace(os.sep, '/')}")
    
    return nav_lines

if __name__ == "__main__":
    notes_dir = "notes"
    if len(sys.argv) > 1:
        notes_dir = sys.argv[1]
    nav_output = build_nav(notes_dir)
    print("nav:")
    for line in nav_output:
        print(line)
