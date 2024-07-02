# Mackerel plugin for apcupsd.

This plugin collects metrics from an APC UPS (Uninterruptible Power Supply) using the `apcaccess` command and reports them to Mackerel.

<img width="388" alt="image" src="https://github.com/matsubo/mackerel-plugin-apcupsd/assets/98103/2b66ec1a-209e-494e-9c61-75a9d255c3ac">

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

3. Place the built binary in a directory where the Mackerel agent can access it (e.g., `/opt/mackerel-agent/plugins/bin/`):
   ```
   sudo mkdir -p /opt/mackerel-agent/plugins/bin/
   sudo cp mackerel-plugin-apcupsd /opt/mackerel-agent/plugins/bin/
   ```

## Configuration

Add the following configuration to your Mackerel agent configuration file (usually located at `/etc/mackerel-agent/mackerel-agent.conf`):

```
[plugin.metrics.apcapcupsd]
command = "/opt/mackerel-agent/plugins/bin/mackerel-plugin-apcapcupsd"
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

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

