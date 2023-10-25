#!/bin/bash
find . ! -name "*.sh" | sed 's/.\///g' | sort -r
