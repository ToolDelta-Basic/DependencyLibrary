local_record=$(sed -n '1,1p' commit)
remote_commit=$(sed -n '1,1p' neomega_core/.git/refs/remotes/origin/main)
# prepare
if [ $local_record != $remote_commit ]; then
    echo "${remote_commit}" > commit
    echo 'changed=true' >> $GITHUB_OUTPUT
    # if changed, we sync the commit id and build depends
else
    echo 'changed=false' >> $GITHUB_OUTPUT
    # if not changed, we do NOTHING
fi
# test commit changes