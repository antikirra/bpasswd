# 🔐 BCrypt Password Generator

This tool is designed to provide an easy and straightforward way to generate BCrypt-protected passwords. Unlike the
`htpasswd` utility, it simplifies the process of creating secure passwords, making it more user-friendly and efficient
for various applications.

## Basic usage

Run:

```console
bpasswd -p 'my$upErP4SsWord'
```

Output:

```console
$2a$12$5gJXBu3/G6lBX3WRurmzi.BLAF16oI5Ip/NWLLolRgc7Uyjfm67ya
```

> [!CAUTION]
> A secure **password should contain special characters** that may affect the input value in the console. Be extremely
> careful and remember to enclose your password in single quotes.

### Pretty result

Sometimes, the presence of extraneous characters like dots and slashes can cause issues when transmitting and storing
the password hash. To prevent their appearance in the output value, use the `-pretty` flag.

Run:

```console
bpasswd -p 'my$upErP4SsWord' -pretty
```

Output:

```console
$2a$12$3IIH2hOBsX3Rjf1HjO7ggesoFPfuszK367R4L1um6nKYVcY9uVg1u
```

> [!CAUTION]
> Attempts to generate a nice hash may take much longer and still not produce the desired result. Use this option with
> extreme caution and at your own risk.

### Don't forget about secure input

Direct input in the console is highly discouraged, as it leaves traces in the system and increases the risk of password
theft. It is strongly recommended to use interactive input to mitigate these risks.

Run:

```console
bpasswd -i
```

By using the `-i` flag, you activate secure password input.

```console
Enter password: 
```

Simply enter the password and press Enter.

### Don't forget about the computational complexity

By default, the `cost` value is set to 12. Currently, this is sufficient to maintain a balance between cache generation
speed and cryptographic strength. However, you can always increase this parameter at your discretion. The most
preferable value would be in the range from 12 to 18.

```console
bpasswd -p 'my$upErP4SsWord' -с 16
```

Output:

```console
$2a$16$fEZET4xQbBVPYLmcmuorSuLGbTHJnvv3bTiMzJF.g4MKM/3apeUye
```

> [!CAUTION]
> A high cost value will lead to an increase in the time required to generate and verify the password hash.

### Output redirection

By default, the result of password hashing is output to the console along with a newline character. This is very
convenient for reading and copying but not ideal for storing in a file when redirecting output. To remove the newline
character, use the `-trim` parameter.

```console
bpasswd -p 'my$upErP4SsWord' -trim > password.txt
```

### Hash verification

To compare the password with the hash, use the `-verify` flag. If successful, the command will return an exit code of
0. If the verification fails, it will return 1.

Run:

```console
bpasswd -p 'my$upErP4SsWord' -verify '$2a$16$fEZET4xQbBVPYLmcmuorSuLGbTHJnvv3bTiMzJF.g4MKM/3apeUye'
```

Output:

```console
Password verification succeeded
```

## Download the Latest Release

To download the latest version of the tool, go to the
GitHub [releases page](https://github.com/antikirra/bpasswd/releases).

Choose the appropriate version and download the archive with the binary file for your operating system.

To download via the command line, use the following command:

```console
curl -sL https://github.com/antikirra/bpasswd/releases/download/v1.5.1/bpasswd_1.5.1_linux_386 -o /usr/bin/bpasswd
```

This will download the 32-bit (386) Linux archive.

Now, you can run `bpasswd` as a regular program anywhere in the system.