#include <errno.h>
#include <iconv.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char *do_iconv(char *to, char *from, char *text)
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

	/* Open conversioon descriptor */
	desc = iconv_open(to, from);
	if ((long)desc == -1) {
		return NULL; /* conversion not supported */
	}

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
