package main

/*
#include <sys/socket.h>
#include <sys/ioctl.h>
#include <linux/if.h>
#include <linux/if_tun.h>

# define _IOW(x,y,z)	(((x)<<8)|y)

#define IFF_TUN		0x0001
#define IFF_TAP		0x0002
#define IFF_NO_PI	0x1000
#define TUNSETIFF     _IOW('T', 202, int) 
#define	IFNAMSIZ	16
*/
//import "C"

const (
	IFF_TUN   = 0x0001
	IFF_TAP   = 0x0002
	IFF_NO_PI = 0x1000
	TUNSETIFF = ((('T')<<8)|202)
)

type ifreq struct {
	name  [16]byte
	flags uint16
	_     [16 - 2]byte
}
