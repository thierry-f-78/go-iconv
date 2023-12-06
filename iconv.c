#include <errno.h>
#include <iconv.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char *do_iconv(char *to, char *from, char *text, int transliterate)
{
	iconv_t desc;
	char *inbuf;
	size_t inbytesleft;
	char *outbuf;
	size_t outbytesleft;
	size_t insize;
	int loop;
	size_t res;
	char *outbuf_start;
	char to_with_transliteration[128];

#ifndef __APPLE__
	if (transliterate) {
		snprintf(to_with_transliteration, 127, "%s//TRANSLIT", to);
		to_with_transliteration[127] = 0;
		to = to_with_transliteration;
	}
#endif

	/* Open conversioon descriptor */
	desc = iconv_open(to, from);
	if ((long)desc == -1) {
		return NULL; /* conversion not supported */
	}

#ifdef __APPLE__
	/* Set option */
	iconvctl(desc, ICONV_SET_TRANSLITERATE, &transliterate);
#endif

	/* Try conversion we reach succes of definitive error */
	loop = 1;
	outbuf_start = NULL;
	insize = strlen(text);
	while (1) {

		/* update output buffers */
		loop++;
		free(outbuf_start);
		outbuf_start = malloc(loop * insize);
		if (outbuf_start == NULL) {
			iconv_close(desc);
			return NULL; /* out of memory error */
		}
		
		/* init conversion vars */
		inbuf = text;
		inbytesleft = insize;
		outbuf = outbuf_start;
		outbytesleft = (loop * insize) - 1; /* ensure one byte for the final \0 */
		res = iconv(desc, &inbuf, &inbytesleft, &outbuf, &outbytesleft);
		if (res == -1) {
			if (errno == E2BIG)
				/* try again with bigger buffer */
				continue;

			/* unrecoverable error */
			iconv_close(desc);
			free(outbuf_start);
			return NULL;
		}
		break;
	}

	/* close converter */
	iconv_close(desc);
	outbuf_start[outbuf - outbuf_start] = '\0';
	return outbuf_start;
}
