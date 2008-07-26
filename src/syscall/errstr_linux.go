// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

export errstr

const (
	ENONE=0;
	EPERM=1;
	ENOENT=2;
	ESRCH=3;
	EINTR=4;
	EIO=5;
	ENXIO=6;
	E2BIG=7;
	ENOEXEC=8;
	EBADF=9;
	ECHILD=10;
	EAGAIN=11;
	ENOMEM=12;
	EACCES=13;
	EFAULT=14;
	ENOTBLK=15;
	EBUSY=16;
	EEXIST=17;
	EXDEV=18;
	ENODEV=19;
	ENOTDIR=20;
	EISDIR=21;
	EINVAL=22;
	ENFILE=23;
	EMFILE=24;
	ENOTTY=25;
	ETXTBSY=26;
	EFBIG=27;
	ENOSPC=28;
	ESPIPE=29;
	EROFS=30;
	EMLINK=31;
	EPIPE=32;
	EDOM=33;
	ERANGE=34;
	EDEADLK=35;
	ENAMETOOLONG=36;
	ENOLCK=37;
	ENOSYS=38;
	ENOTEMPTY=39;
	ELOOP=40;
	ENOMSG=42;
	EIDRM=43;
	ECHRNG=44;
	EL2NSYNC=45;
	EL3HLT=46;
	EL3RST=47;
	ELNRNG=48;
	EUNATCH=49;
	ENOCSI=50;
	EL2HLT=51;
	EBADE=52;
	EBADR=53;
	EXFULL=54;
	ENOANO=55;
	EBADRQC=56;
	EBADSLT=57;
	EBFONT=59;
	ENOSTR=60;
	ENODATA=61;
	ETIME=62;
	ENOSR=63;
	ENONET=64;
	ENOPKG=65;
	EREMOTE=66;
	ENOLINK=67;
	EADV=68;
	ESRMNT=69;
	ECOMM=70;
	EPROTO=71;
	EMULTIHOP=72;
	EDOTDOT=73;
	EBADMSG=74;
	EOVERFLOW=75;
	ENOTUNIQ=76;
	EBADFD=77;
	EREMCHG=78;
	ELIBACC=79;
	ELIBBAD=80;
	ELIBSCN=81;
	ELIBMAX=82;
	ELIBEXEC=83;
	EILSEQ=84;
	ERESTART=85;
	ESTRPIPE=86;
	EUSERS=87;
	ENOTSOCK=88;
	EDESTADDRREQ=89;
	EMSGSIZE=90;
	EPROTOTYPE=91;
	ENOPROTOOPT=92;
	EPROTONOSUPPORT=93;
	ESOCKTNOSUPPORT=94;
	EOPNOTSUPP=95;
	EPFNOSUPPORT=96;
	EAFNOSUPPORT=97;
	EADDRINUSE=98;
	EADDRNOTAVAIL=99;
	ENETDOWN=100;
	ENETUNREACH=101;
	ENETRESET=102;
	ECONNABORTED=103;
	ECONNRESET=104;
	ENOBUFS=105;
	EISCONN=106;
	ENOTCONN=107;
	ESHUTDOWN=108;
	ETOOMANYREFS=109;
	ETIMEDOUT=110;
	ECONNREFUSED=111;
	EHOSTDOWN=112;
	EHOSTUNREACH=113;
	EALREADY=114;
	EINPROGRESS=115;
	ESTALE=116;
	EUCLEAN=117;
	ENOTNAM=118;
	ENAVAIL=119;
	EISNAM=120;
	EREMOTEIO=121;
	EDQUOT=122;
	ENOMEDIUM=123;
	EMEDIUMTYPE=124;
	ECANCELED=125;
	ENOKEY=126;
	EKEYEXPIRED=127;
	EKEYREVOKED=128;
	EKEYREJECTED=129;
	ELAST=130;
)

export (
	ENONE,
	EPERM,
	ENOENT,
	ESRCH,
	EINTR,
	EIO,
	ENXIO,
	E2BIG,
	ENOEXEC,
	EBADF,
	ECHILD,
	EAGAIN,
	ENOMEM,
	EACCES,
	EFAULT,
	ENOTBLK,
	EBUSY,
	EEXIST,
	EXDEV,
	ENODEV,
	ENOTDIR,
	EISDIR,
	EINVAL,
	ENFILE,
	EMFILE,
	ENOTTY,
	ETXTBSY,
	EFBIG,
	ENOSPC,
	ESPIPE,
	EROFS,
	EMLINK,
	EPIPE,
	EDOM,
	ERANGE,
	EDEADLK,
	ENAMETOOLONG,
	ENOLCK,
	ENOSYS,
	ENOTEMPTY,
	ELOOP,
	ENOMSG,
	EIDRM,
	ECHRNG,
	EL2NSYNC,
	EL3HLT,
	EL3RST,
	ELNRNG,
	EUNATCH,
	ENOCSI,
	EL2HLT,
	EBADE,
	EBADR,
	EXFULL,
	ENOANO,
	EBADRQC,
	EBADSLT,
	EBFONT,
	ENOSTR,
	ENODATA,
	ETIME,
	ENOSR,
	ENONET,
	ENOPKG,
	EREMOTE,
	ENOLINK,
	EADV,
	ESRMNT,
	ECOMM,
	EPROTO,
	EMULTIHOP,
	EDOTDOT,
	EBADMSG,
	EOVERFLOW,
	ENOTUNIQ,
	EBADFD,
	EREMCHG,
	ELIBACC,
	ELIBBAD,
	ELIBSCN,
	ELIBMAX,
	ELIBEXEC,
	EILSEQ,
	ERESTART,
	ESTRPIPE,
	EUSERS,
	ENOTSOCK,
	EDESTADDRREQ,
	EMSGSIZE,
	EPROTOTYPE,
	ENOPROTOOPT,
	EPROTONOSUPPORT,
	ESOCKTNOSUPPORT,
	EOPNOTSUPP,
	EPFNOSUPPORT,
	EAFNOSUPPORT,
	EADDRINUSE,
	EADDRNOTAVAIL,
	ENETDOWN,
	ENETUNREACH,
	ENETRESET,
	ECONNABORTED,
	ECONNRESET,
	ENOBUFS,
	EISCONN,
	ENOTCONN,
	ESHUTDOWN,
	ETOOMANYREFS,
	ETIMEDOUT,
	ECONNREFUSED,
	EHOSTDOWN,
	EHOSTUNREACH,
	EALREADY,
	EINPROGRESS,
	ESTALE,
	EUCLEAN,
	ENOTNAM,
	ENAVAIL,
	EISNAM,
	EREMOTEIO,
	EDQUOT,
	ENOMEDIUM,
	EMEDIUMTYPE,
	ECANCELED,
	ENOKEY,
	EKEYEXPIRED,
	EKEYREVOKED,
	EKEYREJECTED,
	ELAST
)

var error [ELAST]string;

func init(){
	error[ENONE] = "No error";
	error[EPERM] = "Operation not permitted";
	error[ENOENT] = "No such file or directory";
	error[ESRCH] = "No such process";
	error[EINTR] = "Interrupted system call";
	error[EIO] = "I/O error";
	error[ENXIO] = "No such device or address";
	error[E2BIG] = "Argument list too long";
	error[ENOEXEC] = "Exec format error";
	error[EBADF] = "Bad file number";
	error[ECHILD] = "No child processes";
	error[EAGAIN] = "Try again";
	error[ENOMEM] = "Out of memory";
	error[EACCES] = "Permission denied";
	error[EFAULT] = "Bad address";
	error[ENOTBLK] = "Block device required";
	error[EBUSY] = "Device or resource busy";
	error[EEXIST] = "File exists";
	error[EXDEV] = "Cross-device link";
	error[ENODEV] = "No such device";
	error[ENOTDIR] = "Not a directory";
	error[EISDIR] = "Is a directory";
	error[EINVAL] = "Invalid argument";
	error[ENFILE] = "File table overflow";
	error[EMFILE] = "Too many open files";
	error[ENOTTY] = "Not a typewriter";
	error[ETXTBSY] = "Text file busy";
	error[EFBIG] = "File too large";
	error[ENOSPC] = "No space left on device";
	error[ESPIPE] = "Illegal seek";
	error[EROFS] = "Read-only file system";
	error[EMLINK] = "Too many links";
	error[EPIPE] = "Broken pipe";
	error[EDOM] = "Math argument out of domain of func";
	error[ERANGE] = "Math result not representable";
	error[EDEADLK] = "Resource deadlock would occur";
	error[ENAMETOOLONG] = "File name too long";
	error[ENOLCK] = "No record locks available";
	error[ENOSYS] = "Function not implemented";
	error[ENOTEMPTY] = "Directory not empty";
	error[ELOOP] = "Too many symbolic links encountered";
	error[ENOMSG] = "No message of desired type";
	error[EIDRM] = "Identifier removed";
	error[ECHRNG] = "Channel number out of range";
	error[EL2NSYNC] = "Level 2 not synchronized";
	error[EL3HLT] = "Level 3 halted";
	error[EL3RST] = "Level 3 reset";
	error[ELNRNG] = "Link number out of range";
	error[EUNATCH] = "Protocol driver not attached";
	error[ENOCSI] = "No CSI structure available";
	error[EL2HLT] = "Level 2 halted";
	error[EBADE] = "Invalid exchange";
	error[EBADR] = "Invalid request descriptor";
	error[EXFULL] = "Exchange full";
	error[ENOANO] = "No anode";
	error[EBADRQC] = "Invalid request code";
	error[EBADSLT] = "Invalid slot";
	error[EBFONT] = "Bad font file format";
	error[ENOSTR] = "Device not a stream";
	error[ENODATA] = "No data available";
	error[ETIME] = "Timer expired";
	error[ENOSR] = "Out of streams resources";
	error[ENONET] = "Machine is not on the network";
	error[ENOPKG] = "Package not installed";
	error[EREMOTE] = "Object is remote";
	error[ENOLINK] = "Link has been severed";
	error[EADV] = "Advertise error";
	error[ESRMNT] = "Srmount error";
	error[ECOMM] = "Communication error on send";
	error[EPROTO] = "Protocol error";
	error[EMULTIHOP] = "Multihop attempted";
	error[EDOTDOT] = "RFS specific error";
	error[EBADMSG] = "Not a data message";
	error[EOVERFLOW] = "Value too large for defined data type";
	error[ENOTUNIQ] = "Name not unique on network";
	error[EBADFD] = "File descriptor in bad state";
	error[EREMCHG] = "Remote address changed";
	error[ELIBACC] = "Can not access a needed shared library";
	error[ELIBBAD] = "Accessing a corrupted shared library";
	error[ELIBSCN] = ".lib section in a.out corrupted";
	error[ELIBMAX] = "Attempting to link in too many shared libraries";
	error[ELIBEXEC] = "Cannot exec a shared library directly";
	error[EILSEQ] = "Illegal byte sequence";
	error[ERESTART] = "Interrupted system call should be restarted";
	error[ESTRPIPE] = "Streams pipe error";
	error[EUSERS] = "Too many users";
	error[ENOTSOCK] = "Socket operation on non-socket";
	error[EDESTADDRREQ] = "Destination address required";
	error[EMSGSIZE] = "Message too long";
	error[EPROTOTYPE] = "Protocol wrong type for socket";
	error[ENOPROTOOPT] = "Protocol not available";
	error[EPROTONOSUPPORT] = "Protocol not supported";
	error[ESOCKTNOSUPPORT] = "Socket type not supported";
	error[EOPNOTSUPP] = "Operation not supported on transport endpoint";
	error[EPFNOSUPPORT] = "Protocol family not supported";
	error[EAFNOSUPPORT] = "Address family not supported by protocol";
	error[EADDRINUSE] = "Address already in use";
	error[EADDRNOTAVAIL] = "Cannot assign requested address";
	error[ENETDOWN] = "Network is down";
	error[ENETUNREACH] = "Network is unreachable";
	error[ENETRESET] = "Network dropped connection because of reset";
	error[ECONNABORTED] = "Software caused connection abort";
	error[ECONNRESET] = "Connection reset by peer";
	error[ENOBUFS] = "No buffer space available";
	error[EISCONN] = "Transport endpoint is already connected";
	error[ENOTCONN] = "Transport endpoint is not connected";
	error[ESHUTDOWN] = "Cannot send after transport endpoint shutdown";
	error[ETOOMANYREFS] = "Too many references: cannot splice";
	error[ETIMEDOUT] = "Connection timed out";
	error[ECONNREFUSED] = "Connection refused";
	error[EHOSTDOWN] = "Host is down";
	error[EHOSTUNREACH] = "No route to host";
	error[EALREADY] = "Operation already in progress";
	error[EINPROGRESS] = "Operation now in progress";
	error[ESTALE] = "Stale NFS file handle";
	error[EUCLEAN] = "Structure needs cleaning";
	error[ENOTNAM] = "Not a XENIX named type file";
	error[ENAVAIL] = "No XENIX semaphores available";
	error[EISNAM] = "Is a named type file";
	error[EREMOTEIO] = "Remote I/O error";
	error[EDQUOT] = "Quota exceeded";
	error[ENOMEDIUM] = "No medium found";
	error[EMEDIUMTYPE] = "Wrong medium type";
	error[ECANCELED] = "Operation Canceled";
	error[ENOKEY] = "Required key not available";
	error[EKEYEXPIRED] = "Key has expired";
	error[EKEYREVOKED] = "Key has been revoked";
	error[EKEYREJECTED] = "Key was rejected by service";
}

var digits string = "0123456789"

func str(val int64) string {  // do it here rather than with fmt to avoid dependency
	if val < 0 {
		return "-" + str(-val);
	}
	var buf [32]byte;  // big enough for int64
	i := len(buf)-1;
	for val >= 10 {
		buf[i] = digits[val%10];
		i--;
		val /= 10;
	}
	buf[i] = digits[val];
	return string(buf)[i:len(buf)];
}

func errstr(errno int64) string {
	if errno < 0 || errno >= len(error) {
		return "Error " + str(errno)
	}
	return error[errno]
}
