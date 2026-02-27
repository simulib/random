
void __assert_rtn(const char *, const char *, int, const char *) __attribute__((__noreturn__)) __attribute__((__cold__)) __attribute__((__disable_tail_calls__));
typedef signed char int8_t;
typedef short int16_t;
typedef int int32_t;
typedef long long int64_t;

typedef unsigned char uint8_t;
typedef unsigned short uint16_t;
typedef unsigned int uint32_t;
typedef unsigned long long uint64_t;
typedef int8_t int_least8_t;
typedef int16_t int_least16_t;
typedef int32_t int_least32_t;
typedef int64_t int_least64_t;
typedef uint8_t uint_least8_t;
typedef uint16_t uint_least16_t;
typedef uint32_t uint_least32_t;
typedef uint64_t uint_least64_t;
typedef int8_t int_fast8_t;
typedef int16_t int_fast16_t;
typedef int32_t int_fast32_t;
typedef int64_t int_fast64_t;
typedef uint8_t uint_fast8_t;
typedef uint16_t uint_fast16_t;
typedef uint32_t uint_fast32_t;
typedef uint64_t uint_fast64_t;
typedef signed char __int8_t;
typedef unsigned char __uint8_t;
typedef short __int16_t;
typedef unsigned short __uint16_t;
typedef int __int32_t;
typedef unsigned int __uint32_t;
typedef long long __int64_t;
typedef unsigned long long __uint64_t;
typedef long __darwin_intptr_t;
typedef unsigned int __darwin_natural_t;
typedef int __darwin_ct_rune_t;
typedef union {
 char __mbstate8[128];
 long long _mbstateL;
} __mbstate_t;
typedef __mbstate_t __darwin_mbstate_t;
typedef long int __darwin_ptrdiff_t;
typedef long unsigned int __darwin_size_t;
typedef __builtin_va_list __darwin_va_list;
typedef int __darwin_wchar_t;
typedef __darwin_wchar_t __darwin_rune_t;
typedef int __darwin_wint_t;
typedef unsigned long __darwin_clock_t;
typedef __uint32_t __darwin_socklen_t;
typedef long __darwin_ssize_t;
typedef long __darwin_time_t;
typedef __int64_t __darwin_blkcnt_t;
typedef __int32_t __darwin_blksize_t;
typedef __int32_t __darwin_dev_t;
typedef unsigned int __darwin_fsblkcnt_t;
typedef unsigned int __darwin_fsfilcnt_t;
typedef __uint32_t __darwin_gid_t;
typedef __uint32_t __darwin_id_t;
typedef __uint64_t __darwin_ino64_t;
typedef __darwin_ino64_t __darwin_ino_t;
typedef __darwin_natural_t __darwin_mach_port_name_t;
typedef __darwin_mach_port_name_t __darwin_mach_port_t;
typedef __uint16_t __darwin_mode_t;
typedef __int64_t __darwin_off_t;
typedef __int32_t __darwin_pid_t;
typedef __uint32_t __darwin_sigset_t;
typedef __int32_t __darwin_suseconds_t;
typedef __uint32_t __darwin_uid_t;
typedef __uint32_t __darwin_useconds_t;
typedef unsigned char __darwin_uuid_t[16];
typedef char __darwin_uuid_string_t[37];
struct __darwin_pthread_handler_rec {
 void (*__routine)(void *);
 void *__arg;
 struct __darwin_pthread_handler_rec *__next;
};
struct _opaque_pthread_attr_t {
 long __sig;
 char __opaque[56];
};
struct _opaque_pthread_cond_t {
 long __sig;
 char __opaque[40];
};
struct _opaque_pthread_condattr_t {
 long __sig;
 char __opaque[8];
};
struct _opaque_pthread_mutex_t {
 long __sig;
 char __opaque[56];
};
struct _opaque_pthread_mutexattr_t {
 long __sig;
 char __opaque[8];
};
struct _opaque_pthread_once_t {
 long __sig;
 char __opaque[8];
};
struct _opaque_pthread_rwlock_t {
 long __sig;
 char __opaque[192];
};
struct _opaque_pthread_rwlockattr_t {
 long __sig;
 char __opaque[16];
};
struct _opaque_pthread_t {
 long __sig;
 struct __darwin_pthread_handler_rec *__cleanup_stack;
 char __opaque[8176];
};
typedef struct _opaque_pthread_attr_t __darwin_pthread_attr_t;
typedef struct _opaque_pthread_cond_t __darwin_pthread_cond_t;
typedef struct _opaque_pthread_condattr_t __darwin_pthread_condattr_t;
typedef unsigned long __darwin_pthread_key_t;
typedef struct _opaque_pthread_mutex_t __darwin_pthread_mutex_t;
typedef struct _opaque_pthread_mutexattr_t __darwin_pthread_mutexattr_t;
typedef struct _opaque_pthread_once_t __darwin_pthread_once_t;
typedef struct _opaque_pthread_rwlock_t __darwin_pthread_rwlock_t;
typedef struct _opaque_pthread_rwlockattr_t __darwin_pthread_rwlockattr_t;
typedef struct _opaque_pthread_t *__darwin_pthread_t;
typedef __darwin_intptr_t intptr_t;
typedef unsigned long uintptr_t;
typedef long int intmax_t;
typedef long unsigned int uintmax_t;
static __inline__ int haveAESNI(){
    return 0;
}
struct r123array1x32{ uint32_t v[1]; };
struct r123array2x32{ uint32_t v[2]; };
struct r123array4x32{ uint32_t v[4]; };
struct r123array8x32{ uint32_t v[8]; };
struct r123array1x64{ uint64_t v[1]; };
struct r123array2x64{ uint64_t v[2]; };
struct r123array4x64{ uint64_t v[4]; };
struct r123array16x8{ uint8_t v[16]; };
 static __inline__ uint32_t mulhilo32(uint32_t a, uint32_t b, uint32_t* hip){ uint64_t product = ((uint64_t)a)*((uint64_t)b); *hip = product>>32; return (uint32_t)product; }
 static __inline__ uint64_t mulhilo64(uint64_t a, uint64_t b, uint64_t* hip){ __uint128_t product = ((__uint128_t)a)*((__uint128_t)b); *hip = product>>64; return (uint64_t)product; }
 static __inline__ struct r123array1x32 _philox2x32bumpkey( struct r123array1x32 key) { key.v[0] += ((uint32_t)0x9E3779B9); return key; }
 static __inline__ struct r123array2x32 _philox4x32bumpkey( struct r123array2x32 key) { key.v[0] += ((uint32_t)0x9E3779B9); key.v[1] += ((uint32_t)0xBB67AE85); return key; }
 static __inline__ struct r123array2x32 _philox2x32round(struct r123array2x32 ctr, struct r123array1x32 key) __attribute__((always_inline)); static __inline__ struct r123array2x32 _philox2x32round(struct r123array2x32 ctr, struct r123array1x32 key){ uint32_t hi; uint32_t lo = mulhilo32(((uint32_t)0xd256d193), ctr.v[0], &hi); struct r123array2x32 out = {{hi^key.v[0]^ctr.v[1], lo}}; return out; }
 static __inline__ struct r123array4x32 _philox4x32round(struct r123array4x32 ctr, struct r123array2x32 key) __attribute__((always_inline)); static __inline__ struct r123array4x32 _philox4x32round(struct r123array4x32 ctr, struct r123array2x32 key){ uint32_t hi0; uint32_t hi1; uint32_t lo0 = mulhilo32(((uint32_t)0xD2511F53), ctr.v[0], &hi0); uint32_t lo1 = mulhilo32(((uint32_t)0xCD9E8D57), ctr.v[2], &hi1); struct r123array4x32 out = {{hi1^ctr.v[1]^key.v[0], lo1, hi0^ctr.v[3]^key.v[1], lo0}}; return out; }
enum r123_enum_philox2x32 { philox2x32_rounds = 10 }; typedef struct r123array2x32 philox2x32_ctr_t; typedef struct r123array1x32 philox2x32_key_t; typedef struct r123array1x32 philox2x32_ukey_t; static __inline__ philox2x32_key_t philox2x32keyinit(philox2x32_ukey_t uk) { return uk; } static __inline__ philox2x32_ctr_t philox2x32_R(unsigned int R, philox2x32_ctr_t ctr, philox2x32_key_t key) __attribute__((always_inline)); static __inline__ philox2x32_ctr_t philox2x32_R(unsigned int R, philox2x32_ctr_t ctr, philox2x32_key_t key) { (__builtin_expect(!(R<=16), 0) ? __assert_rtn(__func__, "philox.h", 349, "R<=16") : (void)0); if(R>0){ ctr = _philox2x32round(ctr, key); } if(R>1){ key = _philox2x32bumpkey(key); ctr = _philox2x32round(ctr, key); } if(R>2){ key = _philox2x32bumpkey(key); ctr = _philox2x32round(ctr, key); } if(R>3){ key = _philox2x32bumpkey(key); ctr = _philox2x32round(ctr, key); } if(R>4){ key = _philox2x32bumpkey(key); ctr = _philox2x32round(ctr, key); } if(R>5){ key = _philox2x32bumpkey(key); ctr = _philox2x32round(ctr, key); } if(R>6){ key = _philox2x32bumpkey(key); ctr = _philox2x32round(ctr, key); } if(R>7){ key = _philox2x32bumpkey(key); ctr = _philox2x32round(ctr, key); } if(R>8){ key = _philox2x32bumpkey(key); ctr = _philox2x32round(ctr, key); } if(R>9){ key = _philox2x32bumpkey(key); ctr = _philox2x32round(ctr, key); } if(R>10){ key = _philox2x32bumpkey(key); ctr = _philox2x32round(ctr, key); } if(R>11){ key = _philox2x32bumpkey(key); ctr = _philox2x32round(ctr, key); } if(R>12){ key = _philox2x32bumpkey(key); ctr = _philox2x32round(ctr, key); } if(R>13){ key = _philox2x32bumpkey(key); ctr = _philox2x32round(ctr, key); } if(R>14){ key = _philox2x32bumpkey(key); ctr = _philox2x32round(ctr, key); } if(R>15){ key = _philox2x32bumpkey(key); ctr = _philox2x32round(ctr, key); } return ctr; }
enum r123_enum_philox4x32 { philox4x32_rounds = 10 }; typedef struct r123array4x32 philox4x32_ctr_t; typedef struct r123array2x32 philox4x32_key_t; typedef struct r123array2x32 philox4x32_ukey_t; static __inline__ philox4x32_key_t philox4x32keyinit(philox4x32_ukey_t uk) { return uk; } static __inline__ philox4x32_ctr_t philox4x32_R(unsigned int R, philox4x32_ctr_t ctr, philox4x32_key_t key) __attribute__((always_inline)); static __inline__ philox4x32_ctr_t philox4x32_R(unsigned int R, philox4x32_ctr_t ctr, philox4x32_key_t key) { (__builtin_expect(!(R<=16), 0) ? __assert_rtn(__func__, "philox.h", 350, "R<=16") : (void)0); if(R>0){ ctr = _philox4x32round(ctr, key); } if(R>1){ key = _philox4x32bumpkey(key); ctr = _philox4x32round(ctr, key); } if(R>2){ key = _philox4x32bumpkey(key); ctr = _philox4x32round(ctr, key); } if(R>3){ key = _philox4x32bumpkey(key); ctr = _philox4x32round(ctr, key); } if(R>4){ key = _philox4x32bumpkey(key); ctr = _philox4x32round(ctr, key); } if(R>5){ key = _philox4x32bumpkey(key); ctr = _philox4x32round(ctr, key); } if(R>6){ key = _philox4x32bumpkey(key); ctr = _philox4x32round(ctr, key); } if(R>7){ key = _philox4x32bumpkey(key); ctr = _philox4x32round(ctr, key); } if(R>8){ key = _philox4x32bumpkey(key); ctr = _philox4x32round(ctr, key); } if(R>9){ key = _philox4x32bumpkey(key); ctr = _philox4x32round(ctr, key); } if(R>10){ key = _philox4x32bumpkey(key); ctr = _philox4x32round(ctr, key); } if(R>11){ key = _philox4x32bumpkey(key); ctr = _philox4x32round(ctr, key); } if(R>12){ key = _philox4x32bumpkey(key); ctr = _philox4x32round(ctr, key); } if(R>13){ key = _philox4x32bumpkey(key); ctr = _philox4x32round(ctr, key); } if(R>14){ key = _philox4x32bumpkey(key); ctr = _philox4x32round(ctr, key); } if(R>15){ key = _philox4x32bumpkey(key); ctr = _philox4x32round(ctr, key); } return ctr; }
 static __inline__ struct r123array1x64 _philox2x64bumpkey( struct r123array1x64 key) { key.v[0] += (0x9E3779B97F4A7C15ULL); return key; }
 static __inline__ struct r123array2x64 _philox4x64bumpkey( struct r123array2x64 key) { key.v[0] += (0x9E3779B97F4A7C15ULL); key.v[1] += (0xBB67AE8584CAA73BULL); return key; }
 static __inline__ struct r123array2x64 _philox2x64round(struct r123array2x64 ctr, struct r123array1x64 key) __attribute__((always_inline)); static __inline__ struct r123array2x64 _philox2x64round(struct r123array2x64 ctr, struct r123array1x64 key){ uint64_t hi; uint64_t lo = mulhilo64((0xD2B74407B1CE6E93ULL), ctr.v[0], &hi); struct r123array2x64 out = {{hi^key.v[0]^ctr.v[1], lo}}; return out; }
 
static __inline__ struct r123array4x64 _philox4x64round(struct r123array4x64 ctr, struct r123array2x64 key) __attribute__((always_inline)); 

static __inline__ struct r123array4x64 _philox4x64round(struct r123array4x64 ctr, struct r123array2x64 key){ 
    uint64_t hi0; 
    uint64_t hi1; 
    uint64_t lo0 = mulhilo64((0xD2E7470EE14C6C93ULL), ctr.v[0], &hi0); 
    uint64_t lo1 = mulhilo64((0xCA5A826395121157ULL), ctr.v[2], &hi1); 
    struct r123array4x64 out = {{hi1^ctr.v[1]^key.v[0], lo1, hi0^ctr.v[3]^key.v[1], lo0}}; 
    return out; 
}


 enum r123_enum_philox2x64 { philox2x64_rounds = 10 }; typedef struct r123array2x64 philox2x64_ctr_t; typedef struct r123array1x64 philox2x64_key_t; typedef struct r123array1x64 philox2x64_ukey_t; static __inline__ philox2x64_key_t philox2x64keyinit(philox2x64_ukey_t uk) { return uk; } static __inline__ philox2x64_ctr_t philox2x64_R(unsigned int R, philox2x64_ctr_t ctr, philox2x64_key_t key) __attribute__((always_inline)); static __inline__ philox2x64_ctr_t philox2x64_R(unsigned int R, philox2x64_ctr_t ctr, philox2x64_key_t key) { (__builtin_expect(!(R<=16), 0) ? __assert_rtn(__func__, "philox.h", 358, "R<=16") : (void)0); if(R>0){ ctr = _philox2x64round(ctr, key); } if(R>1){ key = _philox2x64bumpkey(key); ctr = _philox2x64round(ctr, key); } if(R>2){ key = _philox2x64bumpkey(key); ctr = _philox2x64round(ctr, key); } if(R>3){ key = _philox2x64bumpkey(key); ctr = _philox2x64round(ctr, key); } if(R>4){ key = _philox2x64bumpkey(key); ctr = _philox2x64round(ctr, key); } if(R>5){ key = _philox2x64bumpkey(key); ctr = _philox2x64round(ctr, key); } if(R>6){ key = _philox2x64bumpkey(key); ctr = _philox2x64round(ctr, key); } if(R>7){ key = _philox2x64bumpkey(key); ctr = _philox2x64round(ctr, key); } if(R>8){ key = _philox2x64bumpkey(key); ctr = _philox2x64round(ctr, key); } if(R>9){ key = _philox2x64bumpkey(key); ctr = _philox2x64round(ctr, key); } if(R>10){ key = _philox2x64bumpkey(key); ctr = _philox2x64round(ctr, key); } if(R>11){ key = _philox2x64bumpkey(key); ctr = _philox2x64round(ctr, key); } if(R>12){ key = _philox2x64bumpkey(key); ctr = _philox2x64round(ctr, key); } if(R>13){ key = _philox2x64bumpkey(key); ctr = _philox2x64round(ctr, key); } if(R>14){ key = _philox2x64bumpkey(key); ctr = _philox2x64round(ctr, key); } if(R>15){ key = _philox2x64bumpkey(key); ctr = _philox2x64round(ctr, key); } return ctr; }

 enum r123_enum_philox4x64 { philox4x64_rounds = 10 }; 
 typedef struct r123array4x64 philox4x64_ctr_t; 
 typedef struct r123array2x64 philox4x64_key_t; 
 typedef struct r123array2x64 philox4x64_ukey_t; 
 static __inline__ philox4x64_key_t philox4x64keyinit(philox4x64_ukey_t uk) { return uk; } static __inline__ philox4x64_ctr_t philox4x64_R(unsigned int R, philox4x64_ctr_t ctr, philox4x64_key_t key) __attribute__((always_inline)); static __inline__ philox4x64_ctr_t philox4x64_R(unsigned int R, philox4x64_ctr_t ctr, philox4x64_key_t key) { (__builtin_expect(!(R<=16), 0) ? __assert_rtn(__func__, "philox.h", 359, "R<=16") : (void)0); if(R>0){ ctr = _philox4x64round(ctr, key); } if(R>1){ key = _philox4x64bumpkey(key); ctr = _philox4x64round(ctr, key); } if(R>2){ key = _philox4x64bumpkey(key); ctr = _philox4x64round(ctr, key); } if(R>3){ key = _philox4x64bumpkey(key); ctr = _philox4x64round(ctr, key); } if(R>4){ key = _philox4x64bumpkey(key); ctr = _philox4x64round(ctr, key); } if(R>5){ key = _philox4x64bumpkey(key); ctr = _philox4x64round(ctr, key); } if(R>6){ key = _philox4x64bumpkey(key); ctr = _philox4x64round(ctr, key); } if(R>7){ key = _philox4x64bumpkey(key); ctr = _philox4x64round(ctr, key); } if(R>8){ key = _philox4x64bumpkey(key); ctr = _philox4x64round(ctr, key); } if(R>9){ key = _philox4x64bumpkey(key); ctr = _philox4x64round(ctr, key); } if(R>10){ key = _philox4x64bumpkey(key); ctr = _philox4x64round(ctr, key); } if(R>11){ key = _philox4x64bumpkey(key); ctr = _philox4x64round(ctr, key); } if(R>12){ key = _philox4x64bumpkey(key); ctr = _philox4x64round(ctr, key); } if(R>13){ key = _philox4x64bumpkey(key); ctr = _philox4x64round(ctr, key); } if(R>14){ key = _philox4x64bumpkey(key); ctr = _philox4x64round(ctr, key); } if(R>15){ key = _philox4x64bumpkey(key); ctr = _philox4x64round(ctr, key); } return ctr; }

