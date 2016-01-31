#!/usr/bin/env bash

nx `cat nxlist.txt` | ./gen_rules.pl > rules.txt
