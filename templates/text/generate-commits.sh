{{ define "generate_commits_script" }}
#!/bin/bash
#
########################################################################################
# This script was generated by GitHub Graph Drawer <github-graph-drawer.mbaraa.com>
# Contributer(s): Baraa Al-Masri
########################################################################################
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program.  If not, see <http://www.gnu.org/licenses/>.
########################################################################################
# Run this script on any git repositry (preferbly a new one so you can revert it easily),
# to draw your wanted text "{{ .Message }}" on your GitHub contributions graph.
#
########################################################################################

FILE_NAME=`uuidgen`.txt
MAX_COMMIT_COUNT={{ .CommitsCount }}
DATES=( {{ .Dates }} )

touch $FILE_NAME
for DATE in "${DATES[@]}"; do
    for ((i = 0; i < $MAX_COMMIT_COUNT; i++)); do
        echo `date` >> $FILE_NAME
        git add $FILE_NAME
        git commit -m "🤪 bump commit 😬" --date $DATE
    done
done
{{ end }}
