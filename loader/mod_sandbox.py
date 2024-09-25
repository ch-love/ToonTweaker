import sys

def run_mod_sandboxed(mod_code):
    allowed_globals = {
        'print': print,
        'modify_assets': modify_assets,
        'modify_character': modify_character,
        # Only safe functions are exposed here
    }

    try:
        exec(mod_code, allowed_globals)
    except Exception as e:
        print(f"Mod execution error: {e}")

# Safe API functions
def modify_assets(asset_name, new_value):
    print(f"Modifying asset {asset_name} to {new_value}")

def modify_character(character, new_stats):
    print(f"Modifying character {character} stats: {new_stats}")
