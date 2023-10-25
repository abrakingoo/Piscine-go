#!/bin/bash
directory=$(find . -type d | wc -l)
files=$(find . -type f | wc -l)
echo $(($directory + $files))
