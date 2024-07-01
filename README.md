# Mackerel plugin for apcupsd.

This plugin collects metrics from an APC UPS (Uninterruptible Power Supply) using the `apcaccess` command and reports them to Mackerel.

## Prerequisites

- Go 1.16 or later
- `apcaccess` command (usually provided by the `apcupsd` package)
- Mackerel agent

## Installation

### via mkr command

```
sudo mkr plugin install mackerel-plugin-apcupsd
```

### via GitHub

1. Clone this repository:
   ```
   git clone https://github.com/matsubo/mackerel-plugin-apcupsd.git
   cd mackerel-plugin-apcupsd
   ```

2. Build the plugin:
   ```
   go build -o mackerel-plugin-apcupsd
   ```

3. Place the built binary in a directory where the Mackerel agent can access it (e.g., `/usr/local/bin/`):
   ```
   sudo cp mackerel-plugin-apcupsd /usr/local/bin/
   ```

## Configuration

Add the following configuration to your Mackerel agent configuration file (usually located at `/etc/mackerel-agent/mackerel-agent.conf`):

```
[plugin.metrics.apcapcupsd]
command = "/usr/local/bin/mackerel-plugin-apcapcupsd"
```

## Collected Metrics

This plugin collects the following metrics:

- `custom.ups.linev`: Input line voltage (Volts)
- `custom.ups.loadpct`: UPS load (Percent)
- `custom.ups.bcharge`: Battery charge (Percent)
- `custom.ups.timeleft`: Estimated runtime left (Minutes)
- `custom.ups.battv`: Battery voltage (Volts)

## Testing

To run the tests for this plugin, use the following command:

```
go test -v
```

Note that the `TestFetchMetrics` test requires the `apcaccess` command to be available and will be skipped if it's not.

## Troubleshooting

- Ensure that the `apcupsd` service is running and that the `apcaccess` command is available and executable by the Mackerel agent.
- Check the Mackerel agent logs for any error messages related to this plugin.
- Verify that the UPS is properly connected and recognized by the system.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

