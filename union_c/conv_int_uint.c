#include <stdio.h>
#include <stdlib.h>
#include <limits.h>

typedef struct		n {
	int					type;
	union {
		int				i;
		unsigned		u; 
	}	value;
}					number;

int	main(void)
{
	int		x;
	number	*num;

	num = (number *)malloc(sizeof(number));
	x = -1;
	num->value.i = x;
	printf("test %u == %u\n", num->value.u, UINT_MAX);
	free(num);
}
