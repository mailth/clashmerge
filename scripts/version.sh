#!/bin/bash

set -e

ARGS="$1"
VERSION_VAR="$2"
CURRENT_VERSION=$(cat VERSION)

show_usage() {
    echo "Current version: $CURRENT_VERSION"
    echo "Usage:"
    echo "  make version VERSION=v1.0.0    # Set specific version"
    echo "  make version v1.0.0            # Set specific version"
    echo "  make version patch             # Increment patch version"
    echo "  make version minor             # Increment minor version"
    echo "  make version major             # Increment major version"
    echo "  make version pre               # Add/increment pre-release (default: rc)"
    echo "  make version pre alpha         # Add alpha pre-release"
    echo "  make version pre beta          # Add beta pre-release"
    echo "  make version pre rc            # Add rc pre-release"
    echo "  make version pre alpha 2       # Add specific alpha version"
}

parse_version() {
    local version="$1"
    # Remove 'v' prefix and split by '-'
    local clean_version="${version#v}"
    local base_version="${clean_version%-*}"
    local pre_release=""
    
    if [[ "$clean_version" == *"-"* ]]; then
        pre_release="${clean_version#*-}"
    fi
    
    # Split base version
    IFS='.' read -r major minor patch <<< "$base_version"
    
    echo "$major $minor $patch $pre_release"
}

increment_version() {
    local type="$1"
    local version_parts
    read -r major minor patch pre_release <<< "$(parse_version "$CURRENT_VERSION")"
    
    case "$type" in
        "patch")
            patch=$((patch + 1))
            echo "v$major.$minor.$patch"
            ;;
        "minor")
            minor=$((minor + 1))
            patch=0
            echo "v$major.$minor.$patch"
            ;;
        "major")
            major=$((major + 1))
            minor=0
            patch=0
            echo "v$major.$minor.$patch"
            ;;
    esac
}

handle_prerelease() {
    local args=($1)
    local version_parts
    read -r major minor patch pre_release <<< "$(parse_version "$CURRENT_VERSION")"
    
    local base_version="$major.$minor.$patch"
    local pre_type="rc"
    local pre_num="1"
    
    # Parse arguments
    if [[ ${#args[@]} -ge 3 ]]; then
        # make version pre alpha 2 (明确指定版本号)
        pre_type="${args[1]}"
        pre_num="${args[2]}"
    elif [[ ${#args[@]} -ge 2 ]]; then
        # make version pre alpha (指定类型，智能处理版本号)
        pre_type="${args[1]}"
        if [[ -n "$pre_release" ]]; then
            # 当前已经是 pre-release
            IFS='.' read -r current_type current_num <<< "$pre_release"
            if [[ "$current_type" == "$pre_type" ]]; then
                # 相同类型，递增版本号
                pre_num=$((current_num + 1))
            else
                # 不同类型，重置为 1
                pre_num="1"
            fi
        else
            # 当前是正式版本，从 1 开始
            pre_num="1"
        fi
    else
        # make version pre (智能处理)
        if [[ -n "$pre_release" ]]; then
            # 当前已经是 pre-release，递增版本号
            IFS='.' read -r current_type current_num <<< "$pre_release"
            pre_type="$current_type"
            pre_num=$((current_num + 1))
        else
            # 当前是正式版本，添加默认 pre-release
            pre_type="rc"
            pre_num="1"
        fi
    fi
    
    echo "v$base_version-$pre_type.$pre_num"
}

update_version() {
    local new_version="$1"
    echo "Updating version from $CURRENT_VERSION to $new_version"
    echo "$new_version" > VERSION
    git add .
    git commit -m "update version to $new_version"
    git push origin HEAD
    echo "[$new_version] Version file updated and pushed successfully"
}

main() {
    if [[ -n "$ARGS" ]]; then
        local first_arg="${ARGS%% *}"
        
        case "$first_arg" in
            "patch"|"minor"|"major")
                new_version=$(increment_version "$first_arg")
                ;;
            "pre")
                new_version=$(handle_prerelease "$ARGS")
                ;;
            *)
                new_version="$first_arg"
                ;;
        esac
        
        update_version "$new_version"
    elif [[ -n "$VERSION_VAR" ]]; then
        update_version "$VERSION_VAR"
    else
        show_usage
    fi
}

main "$@"
