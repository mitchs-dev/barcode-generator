# Barcode Generator (QR Code)

barcode-generator is a simple CLI barcode generator that generates a QR code and saves it, as a PNG file, to a specified directory.

For my purposes, I wanted to generate QR codes for my barcodes. If you'd like to explore other barcode types check out [boombuler/barcode](https://github.com/boombuler/barcode) which is the library used to generate the barcodes.

This tool also uses the `addLabel()` function which adds a label to the barcode. This is useful for adding a human-readable label to the barcode, in the event the barcode is damaged, or unreadable for some reason.


## Usage

> **Note**: The input given will be trimmed of any whitespace.

```bash
barcode-generator <input> <output-directory>
```

### Example

```bash
barcode-generator "ABC-123456" .
```

This will generate a barcode with the text "ABC-123456" and save it to `./barcode-ABC-123456.png`:

![barcode-ABC-123456](./barcode-ABC-123456.png)

### Multiple Barcodes

You can also provide a comma-separated list of barcodes to generate multiple barcodes at once.

```bash
barcode-generator "ABC-123456,DEF-789012,GHI-345678" .
```
