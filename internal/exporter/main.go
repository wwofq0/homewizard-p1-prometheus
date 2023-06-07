package exporter

import (
	// "github.com/chrisdoc/homewizard-p1-prometheus/internal/homewizard"
	"github.com/chrisdoc/homewizard-p1-prometheus/internal/homewizard"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	wifiStrength = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "wifi_strength",
		Help: "Wifi strength in Db",
	})
	totalPowerImportKwh = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "total_power_import_kwh",
		Help: "The total power import in kWh",
	})
	totalPowerImportT1Kwh = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "total_power_import_t1_kwh",
		Help: "The total power import on T1 in kWh",
	})
	totalPowerImportT2Kwh = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "total_power_import_t2_kwh",
		Help: "The total power import on T2 in kWh",
	})
	totalPowerExportKwh = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "total_power_export_kwh",
		Help: "The total power export in kWh",
	})
	totalPowerExportT1Kwh = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "total_power_export_t1_kwh",
		Help: "The total power export on T1 in kWh",
	})
	totalPowerExportT2Kwh = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "total_power_export_t2_kwh",
		Help: "The total power export on T2 in kWh",
	})
	activePowerW = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "active_power_w",
		Help: "The active power in Watt",
	})
	activePowerL1W = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "active_power_l1_w",
		Help: "The active power on L1 in Watt",
	})
	activePowerL2W = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "active_power_l2_w",
		Help: "he active power on L2 in Watt",
	})
	activePowerL3W = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "active_power_l3_w",
		Help: "The active power on L3 in Watt",
	})
	activeVoltageL1V = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "active_voltage_l1_v",
		Help: "The active voltage on L1 in Volt",
	})
	activeVoltageL2V = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "active_voltage_l2_v",
		Help: "The active voltage on L2 in Volt",
	})
	activeVoltageL3V = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "active_voltage_l3_v",
		Help: "The active voltage on L3 in Volt",
	})
	activeCurrentL1A = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "active_current_l1_a",
		Help: "The active current on L1 in Amper",
	})
	activeCurrentL2A = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "active_current_l2_a",
		Help: "The active current on L2 in Amper",
	})
	activeCurrentL3A = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "active_current_l3_a",
		Help: "The active current on L3 in Amper",
	})
	activeFrequenczyHZ = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "active_frequency_hz",
		Help: "The active frequenczy",
	})
	totalGasM3 = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "total_gas_m3",
		Help: "The total gas consumption in m3",
	})
)

// Prometheus exporter
type Prometheus struct{}

// SetData for prometheus exporter
func (p *Prometheus) SetData(home *homewizard.Data) {
	wifiStrength.Set(home.WifiStrength)

	totalPowerImportKwh.Set(home.TotalPowerImportKwh)
	totalPowerImportT1Kwh.Set(home.TotalPowerImportT1Kwh)
	totalPowerImportT2Kwh.Set(home.TotalPowerImportT2Kwh)
	totalPowerExportKwh.Set(home.TotalPowerExportKwh)
	totalPowerExportT1Kwh.Set(home.TotalPowerExportT2Kwh)
	totalPowerExportT2Kwh.Set(home.TotalPowerExportT2Kwh)

	activePowerW.Set(home.ActivePowerW)

	activePowerL1W.Set(home.ActivePowerL1W)
	activePowerL2W.Set(home.ActivePowerL2W)
	activePowerL3W.Set(home.ActivePowerL3W)

	activeVoltageL1V.Set(home.ActiveVoltageL1V)
	activeVoltageL2V.Set(home.ActiveVoltageL2V)
	activeVoltageL3V.Set(home.ActiveVoltageL3V)

	activeCurrentL1A.Set(home.ActiveCurrentL1A)
	activeCurrentL2A.Set(home.ActiveCurrentL2A)
	activeCurrentL3A.Set(home.ActiveCurrentL3A)

	activeFrequenczyHZ.Set(home.ActiveFrequenczyHZ)

	totalGasM3.Set(home.TotalGasM3)
}
