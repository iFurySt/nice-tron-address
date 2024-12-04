# Tron Address Generator

This tool generates Tron addresses with a specific suffix. Like:
```plaintext
TBvjAaMRHwvzFbtM3GvxC9mg8888888888
```

## Quickstart
Download the latest release from the [releases page](https://github.com/iFurySt/nice-tron-address/releases).

Run the binary with the desired parameters

### Generate a single address with suffix `888`:
```bash
./tron-addr-gen --suffix 888 --num 1
```

### Generate 5 addresses with suffix `abc`:
```bash
./tron-addr-gen -s abc -n 5
```

### Output Example
```plaintext
Private Key: e3d9b4c8f2f5a...bcdef, Address: TBvjAaMRHwvzFbtM3GvxC9mg8888888888
```

## License
This project is licensed under the **Creative Commons Legal Code CC0 1.0 Universal**. For more information, refer to the [LICENSE](LICENSE) file.
