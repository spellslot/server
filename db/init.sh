#!/bin/bash

psql -d $1 -U $2 -f db/create_spells_table.sql
psql -d $1 -U $2 -f db/create_classes_table.sql
