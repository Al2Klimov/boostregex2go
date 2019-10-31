#include "regex.hpp"
// Regex

#include <boost/regex.hpp>
using boost::bad_expression;

#include <stdint.h>
// uint64_t

#include <utility>
using std::move;

extern "C" unsigned char CompileRegex(uint64_t pattern_start, uint64_t pattern_end, uint64_t out)
{
	try {
		*(Regex<char>*)out = Regex<char>((const char*)pattern_start, (const char*)pattern_end);
	} catch (const boost::bad_expression&) {
		return 2;
	} catch (...) {
		return 1;
	}

	return 0;
}

extern "C" void FreeRegex(uint64_t rgx)
{
	try {
		Regex<char> r (move(*(Regex<char>*)rgx));
	} catch (...) {
	}
}

extern "C" signed char MatchesSomewhere(uint64_t rgx, uint64_t subject_start, uint64_t subject_end)
{
	try {
		return ((const Regex<char>*)rgx)->MatchesSomewhere((const char*)subject_start, (const char*)subject_end);
	} catch (...) {
		return -1;
	}
}
