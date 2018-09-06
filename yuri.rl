// Copyright (c) 2018, Hayden Eskriett <hayden@eskriett.com>
// See LICENSE for licensing details
//
// Build with:
//   ragel -G2 -Z yuri.rl

// Package yuri provides a blazing fast way to extracts URIs from plain text
package yuri

// YankURIs extracts URIs from text
func YankURIs(data []byte) []string {

    cs, p, pe, eof := 0, 0, len(data), len(data)
    ts, te, act := 0, 0, 0
    _ = eof
    _ = ts
    _ = te
    _ = act

    var uris []string

    %%{
        machine yuri;
        write data;

        action addURI { uris = append(uris, string(data[ts:te])) }

        # Loosely based on RFC3987

        pctencoded   = "%" xdigit{2};
        h16          = xdigit{1,4};
        port         = digit{,5};
        decoctet     = ( digit | [1-9] digit | "1" digit{2} | "2" [0-4] digit | "25" [0-5] );
        subdelims    = ( "!" | "$" | "&" | "'" | "(" | ")" | "*" | "+" | "," | ";" | "=" );
        gendelims    = ( ":" | "/" | "?" | "#" | "[" | "]" | "@" );
        unreserved   = ( alnum | "-" | "." | "_" | "~" );
        reserved     = ( gendelims | subdelims );

        ipv4address  = ( decoctet "." ){3} decoctet;
        ls32         = ( h16 ":" h16 ) | ipv4address;
        ipv6address  = (                            (( h16 ":" ){6} ls32))
                     | (                       "::" (( h16 ":" ){5} ls32))
                     | ((                h16)? "::" (( h16 ":" ){4} ls32))
                     | ((( h16 ":" ){,1} h16)? "::" (( h16 ":" ){3} ls32))
                     | ((( h16 ":" ){,2} h16)? "::" (( h16 ":" ){2} ls32))
                     | ((( h16 ":" ){,3} h16)? "::" (( h16 ":" ){1} ls32))
                     | ((( h16 ":" ){,4} h16)? "::" (               ls32))
                     | ((( h16 ":" ){,5} h16)? "::" (               h16 ))
                     | ((( h16 ":" ){,6} h16)? "::" );
        ipvfuture    = "v" xdigit+ "." ( unreserved | subdelims | ":" )+;
        ipliteral    = "[" ( ipv6address | ipvfuture )"]";

        # ucschars and iprivate are massive simplifications (i.e. they will
        # capture any non-ascii character). I may change this in future, but for
        # now it helps simplify the state machine
        ucschar       = ( any -- ascii );
        iprivate      = ucschar;

        alnumucschars = ( alnum | ucschar );
        iunreserved   = ( alnumucschars | "-" | "." | "_" | "~" );

        iregname      = alnumucschars (iunreserved+ alnumucschars)?;
        iuserinfo     = ( iunreserved | pctencoded | subdelims | ":" )*;
        ihost         = ( ipliteral | ipv4address | iregname );
        iauthority    = ( iuserinfo "@" )? ihost ( ":" port )?;

        ipchar        = ( iunreserved | pctencoded | subdelims | ":" | "@" );
        isegment      = ipchar*;
        iquery        = ( ipchar | iprivate | "/" | "?" )*;
        ifragment     = ( ipchar | "/" | "?" )*;

        noschemeiri   = iauthority ( ( "/" isegment )+ ( "?" iquery )? ( "#" ifragment )? )?;

        ftp          = /ftp/i "://";
        http         = /h/i (/tt/i | /xx/i ) /p/i (/s/i)? "://";
        mailto       = /mailto/i ":";

        iri          = ( ftp | http | mailto )  noschemeiri;

        main := |*
            iri => addURI;
            any;
        *|;
    }%%

    %% write init;
    %% write exec;

    return uris
}
