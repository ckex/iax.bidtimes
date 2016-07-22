#!/bin/bash

# scp -P 40022 dsp.iclick.tar.gz ckex@52.68.216.39:/home/ckex/ssp/

# scp dsp.iclick.tar.gz iax-rtb1:/home/ckex/dsp-client/source
scp ../target/linux/*.tar.gz iax-data1:/home/ckex/sources
