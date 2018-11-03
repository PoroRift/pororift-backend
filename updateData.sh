#!/usr/bin/env bash

# Get content version status
curl https://ddragon.leagueoflegends.com/realms/na.json > response

WORKDIR=${PWD}/data
rm -rf $WORKDIR
mkdir -p $WORKDIR

# Parse Version from the response
python << EOF > version

import json 

with open('response') as f:
  data = json.load(f)

  print data["dd"]

EOF


VERSION=$(cat version)


curl https://ddragon.leagueoflegends.com/cdn/dragontail-$VERSION.tgz > $WORKDIR/data.tgz

tar -xvzf $WORKDIR/data.tgz -C $WORKDIR/

mv $WORKDIR/$VERSION $WORKDIR/content

# Remove unused files
rm version
rm response
rm data/data.tgz

