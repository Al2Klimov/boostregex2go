#pragma once

#include <boost/regex.hpp>
// boost::basic_regex
// boost::match_results
// boost::regex_search

#include <utility>
// std::forward
// std::move

template<class Char>
struct Regex
{
	template<class... Args>
	inline
	Regex(Args&&... args) : Rgx(new boost::basic_regex<Char>(std::forward<Args>(args)...))
	{
	}

	Regex(const Regex& origin) : Rgx(new boost::basic_regex<Char>(*origin.Rgx))
	{
	}

	Regex& operator=(const Regex& origin)
	{
		Regex copy (origin);
		return *this = std::move(copy);
	}

	inline
	Regex(Regex&& origin) noexcept : Rgx(origin.Rgx)
	{
		origin.Rgx = nullptr;
	}

	inline
	Regex& operator=(Regex&& origin) noexcept
	{
		this->~Regex();
		new(this) Regex(std::move(origin));
		return *this;
	}

	inline
	~Regex()
	{
		delete this->Rgx;
	}

	template<class Iterator>
	bool MatchesSomewhere(Iterator first, Iterator last) const
	{
		boost::match_results<Iterator> m;
		return boost::regex_search(first, last, m, *(const boost::basic_regex<Char>*)this->Rgx);
	}

	boost::basic_regex<Char>* Rgx;
};
