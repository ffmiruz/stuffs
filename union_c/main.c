#include <stdio.h>
#include <stdarg.h>
#include <unistd.h>
#include <stdlib.h>

typedef struct		s_par {
	int			type;
	const char	*fmt;
	union {
		long long 			i;
		unsigned long long	u; 
	}	value;
}					t_par;


int	ft_putchar(int c)
{
	static int	count;

	write(1, &c, 1);
	return (++count);
}

int	print_u(unsigned int num, int base)
{
	unsigned int	a;
	int				rem;

	a = num / base;
	if (a)
		print_u(a, base);
	rem = num % base + '0';
	return (ft_putchar(rem));
}

int	print_i(long long num, int base)
{
	long long	a;
	int			rem;

	if (num < 0)
	{
		ft_putchar('-');
		num = -num;
	}
	a = num / base;
	if (a)
		print_i(a, base);
	rem = num % base + '0';
	return (ft_putchar(rem));
}

void	print_ctr(t_par *param)
{
	if (param->type == 1)
	{
		print_i(param->value.i, 10);
	}
	if (param->type == 2)
	{
		print_u(param->value.u, 10);
	}
}

t_par *parse_param(t_par *param, va_list args)
{
	int	c;

	c = *(param->fmt)++;
	if (c == 'd' || c == 'i')
	{
		param->value.i = va_arg(args, int);
		param->type = 1;
	}
	if (c == 'u')
	{
		param->value.u = va_arg(args, unsigned int);
		param->type = 2;
	}
	return param;
}

int	ft_printf(const char *fmt, ...)
{
	va_list	args;
	int		count;
	t_par	*param;

	param = (t_par *)malloc(sizeof(t_par));
	param->fmt = fmt;
	count = 0;
	va_start(args, fmt);
	while (1)
	{	
		while (*(param->fmt) != '%')
		{
			if (*(param->fmt) == '\0')
				return (count);
			count = ft_putchar(*(param->fmt)++);
		}
		*++(param->fmt);
		parse_param(param, args);
		print_ctr(param);
	}
	va_end(args);
	free(param);
	return (0);
}

int	main(void)
{
	unsigned int	i;
	int				x;

	i = -2147483648;
	x = ft_printf("test %i %u\n", i, (unsigned int)909);
	printf("write count std %i\n", x);
}