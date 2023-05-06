package __

// directives has code to test *some* build directives (not all could be tested in a simple way)
// They all appear as "comments" (so are not part of the language) of the form //go:<keyword> ...
// Here are the values for <keyword> I know about: build, embed, generate, norace, noinline, noescape
// They may have other restrictions - for example: //go:build must appear at the top of the source
// file and be followed by a blank line and only be used once (per file).
