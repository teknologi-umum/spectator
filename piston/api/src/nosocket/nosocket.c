/*
nosocket.c

Disables access to the `socket` syscall and runs a program provided as the first
commandline argument.
*/
#include <stdio.h>
#include <errno.h>
#include <unistd.h>
#include <sys/prctl.h>
#include <seccomp.h>

int main(int argc, char *argv[])
{
    // Disallow any new capabilities from being added
    prctl(PR_SET_NO_NEW_PRIVS, 1, 0, 0, 0);

    // SCMP_ACT_ALLOW lets the filter have no effect on syscalls not matching a
    // configured filter rule (allow all by default)
    scmp_filter_ctx ctx = seccomp_init(SCMP_ACT_ALLOW);
    if (!ctx)
    {
        fprintf(stderr, "Unable to initialize seccomp filter context\n");
        return 1;
    }

    // Add 32 bit and 64 bit architectures to seccomp filter
    int rc;
    uint32_t arch[] = {SCMP_ARCH_X86_64, SCMP_ARCH_X86, SCMP_ARCH_X32};
    // We first remove the existing arch, otherwise our subsequent call to add
    // it will fail
    seccomp_arch_remove(ctx, seccomp_arch_native());
    for (int i = 0; i < sizeof(arch) / sizeof(arch[0]); i++)
    {
        rc = seccomp_arch_add(ctx, arch[i]);
        if (rc != 0)
        {
            fprintf(stderr, "Unable to add arch: %d\n", arch[i]);
            return 1;
        }
    }

    // Add a seccomp rule to the syscall blacklist - blacklist the socket syscall
    if (seccomp_rule_add(ctx, SCMP_ACT_ERRNO(EACCES), SCMP_SYS(socket), 0) < 0)
    {
        fprintf(stderr, "Unable to add seccomp rule to context\n");
        return 1;
    }

#ifdef DEBUG
    seccomp_export_pfc(ctx, 0);
#endif

    if (argc < 2)
    {
        fprintf(stderr, "Usage %s: %s <program name> <arguments>\n", argv[0], argv[0]);
        return 1;
    }
    seccomp_load(ctx);
    execvp(argv[1], argv + 1);
    return 1;
}
