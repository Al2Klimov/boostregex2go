#pragma once

#include <stdint.h>
// uint64_t

unsigned char CompileRegex(uint64_t pattern_start, uint64_t pattern_end, uint64_t out);

void FreeRegex(uint64_t rgx);

signed char MatchesSomewhere(uint64_t rgx, uint64_t subject_start, uint64_t subject_end);
