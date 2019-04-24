# Go-RC5-32/12/16

RC5 is a symetric-key block cipher designed in 1994. The RC5 is basically denoted as RC5-w/r/b where w=word size in bits, r=number of rounds, b=number of 8-bit bytes in the key. My implementation would be RC5-32/12/16

The encryption and decryption routines can be specified in a few lines of code. The key schedule, however, is more complex, expanding the key using an essentially one-way function with the binary expansions of both e and the golden ratio.

## Algorithm

### Key Expansion
RC5 encryption and decryption both expand the random key into 2(r+1) words that will be used sequentially (and only once each) during the encryption and decryption processes.

#### Step-1: Convert secret key from bytes to words
This step converts the b-byte (b=number of 8-bit bytes in the key) key into a sequence of words stored in the array L.
```
for(i = b-1, L[c-1] = 0; i != -1; i--)
      L[i/u] = (L[i/u] << 8) + K[i];
```

#### Step-2: Initialize the expanded key table
This step fills in the S table with magic constant Pw and Qw.
* For w = 16: Pw = 0xB7E1, Qw = 0x9E37
* For w = 16: Pw = 0xB7E15163, Qw = 0x9E3779B9
* For w = 16: Pw = 0xB7E151628AED2A6B, Qw = 0x9E3779B97F4A7C15
```
for(S[0] = P, i = 1; i < t; i++)
      S[i] = S[i-1] + Q;
```

#### Step-3: Mix in the secret key
This step is mixing secret key L with key table S
```
for(A = B = i = j = k = 0; k < 3 * t; k++, i = (i+1) % t, j = (j+1) % c)
   {
      A = S[i] = ROTL(S[i] + (A + B), 3);
      B = L[j] = ROTL(L[j] + (A + B), (A + B));
   }
```
> t = 2*(r+1), c = number words in key = ceil(8*b/w)
//pic

### Encryption


### Decryption
