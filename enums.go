package go_modemmanager

// todo errors ?  https://gitlab.freedesktop.org/mobile-broadband/ModemManager/-/blob/master/include/ModemManager-errors.h

// ref https://gitlab.freedesktop.org/mobile-broadband/ModemManager/-/blob/master/include/ModemManager-enums.h
type MMModemCapability uint32 //  Flags describing one or more of the general access technology families that a modem supports.

//go:generate stringer -type=MMModemCapability
const (
	MmModemCapabilityNone        MMModemCapability = 0          // Modem has no capabilities.
	MmModemCapabilityPots        MMModemCapability = 1 << 0     //  Modem supports the analog wired telephone network (ie 56k dialup) and does not have wireless/cellular capabilities.
	MmModemCapabilityCdmaEvdo    MMModemCapability = 1 << 1     // Modem supports at least one of CDMA 1xRTT, EVDO revision 0, EVDO revision A, or EVDO revision B.
	MmModemCapabilityGsmUmts     MMModemCapability = 1 << 2     // Modem supports at least one of GSM, GPRS, EDGE, UMTS, HSDPA, HSUPA, or HSPA+ packet switched data capability.
	MmModemCapabilityLte         MMModemCapability = 1 << 3     // Modem has LTE data capability.
	MmModemCapabilityLteAdvanced MMModemCapability = 1 << 4     // Modem has LTE Advanced data capability.
	MmModemCapabilityIridium     MMModemCapability = 1 << 5     //Modem has Iridium capabilities.
	MmModemCapabilityAny         MMModemCapability = 0xFFFFFFFF // Mask specifying all capabilities.
)

type MMModemLock uint32 // Possible lock reasons.
//go:generate stringer -type=MMModemLock
const (
	MmModemLockUnknown     MMModemLock = 0  // Lock reason unknown.
	MmModemLockNone        MMModemLock = 1  // Modem is unlocked.
	MmModemLockSimPin      MMModemLock = 2  // SIM requires the PIN code.
	MmModemLockSimPin2     MMModemLock = 3  // SIM requires the PIN2 code.
	MmModemLockSimPuk      MMModemLock = 4  // SIM requires the PUK code.
	MmModemLockSimPuk2     MMModemLock = 5  // SIM requires the PUK2 code.
	MmModemLockPhSpPin     MMModemLock = 6  // Modem requires the service provider PIN code.
	MmModemLockPhSpPuk     MMModemLock = 7  // Modem requires the service provider PUK code.
	MmModemLockPhNetPin    MMModemLock = 8  // Modem requires the network PIN code.
	MmModemLockPhNetPuk    MMModemLock = 9  // Modem requires the network PUK code.
	MmModemLockPhSimPin    MMModemLock = 10 // Modem requires the PIN code.
	MmModemLockPhCorpPin   MMModemLock = 11 // Modem requires the corporate PIN code.
	MmModemLockPhCorpPuk   MMModemLock = 12 // Modem requires the corporate PUK code.
	MmModemLockPhFsimPin   MMModemLock = 13 // Modem requires the PH-FSIM PIN code.
	MmModemLockPhFsimPuk   MMModemLock = 14 // Modem requires the PH-FSIM PUK code.
	MmModemLockPhNetsubPin MMModemLock = 15 // Modem requires the network subset PIN code.
	MmModemLockPhNetsubPuk MMModemLock = 16 // Modem requires the network subset PUK code.
)

type MMModemState int32 // Possible modem states.
//go:generate stringer -type=MMModemState
const (
	MmModemStateFailed        MMModemState = -1 // The modem is unusable.
	MmModemStateUnknown       MMModemState = 0  // State unknown or not reportable.
	MmModemStateInitializing  MMModemState = 1  // The modem is currently being initialized.
	MmModemStateLocked        MMModemState = 2  // The modem needs to be unlocked.
	MmModemStateDisabled      MMModemState = 3  // The modem is not enabled and is powered down.
	MmModemStateDisabling     MMModemState = 4  // The modem is currently transitioning to the @MmModemStateDisabled state.
	MmModemStateEnabling      MMModemState = 5  // The modem is currently transitioning to the @MmModemStateEnabled state.
	MmModemStateEnabled       MMModemState = 6  // The modem is enabled and powered on but not registered with a network provider and not available for data connections.
	MmModemStateSearching     MMModemState = 7  // The modem is searching for a network provider to register with.
	MmModemStateRegistered    MMModemState = 8  // The modem is registered with a network provider, and data connections and messaging may be available for use.
	MmModemStateDisconnecting MMModemState = 9  // The modem is disconnecting and deactivating the last active packet data bearer. This state will not be entered if more than one packet data bearer is active and one of the active bearers is deactivated.
	MmModemStateConnecting    MMModemState = 10 // The modem is activating and connecting the first packet data bearer. Subsequent bearer activations when another bearer is already active do not cause this state to be entered.
	MmModemStateConnected     MMModemState = 11 // One or more packet data bearers is active and connected.

)

type MMModemStateFailedReason uint32 // Power state of the modem.
//go:generate stringer -type=MMModemStateFailedReason
const (
	MmModemPowerStateUnknown MMModemStateFailedReason = 0 // Unknown power state.
	MmModemPowerStateOff     MMModemStateFailedReason = 1 // Off.
	MmModemPowerStateLow     MMModemStateFailedReason = 2 // Low-power mode.
	MmModemPowerStateOn      MMModemStateFailedReason = 3 // Full power mode.

)

type MMModemStateChangeReason uint32 // Possible reasons to have changed the modem state.
//go:generate stringer -type=MMModemStateChangeReason
const (
	MmModemStateChangeReasonUnknown       MMModemStateChangeReason = 0 // Reason unknown or not reportable.
	MmModemStateChangeReasonUserRequested MMModemStateChangeReason = 1 // State change was requested by an interface user.
	MmModemStateChangeReasonSuspend       MMModemStateChangeReason = 2 // State change was caused by a system suspend.
	MmModemStateChangeReasonFailure       MMModemStateChangeReason = 3 // State change was caused by an unrecoverable error.

)

type MMModemAccessTechnology uint32 // Describes various access technologies that a device uses when registered with or connected to a network.

//go:generate stringer -type=MMModemAccessTechnology
const (
	MmModemAccessTechnologyUnknown    MMModemAccessTechnology = 0          // The access technology used is unknown.
	MmModemAccessTechnologyPots       MMModemAccessTechnology = 1 << 0     // Analog wireline telephone.
	MmModemAccessTechnologyGsm        MMModemAccessTechnology = 1 << 1     // GSM.
	MmModemAccessTechnologyGsmCompact MMModemAccessTechnology = 1 << 2     // Compact GSM.
	MmModemAccessTechnologyGprs       MMModemAccessTechnology = 1 << 3     // GPRS.
	MmModemAccessTechnologyEdge       MMModemAccessTechnology = 1 << 4     // EDGE (ETSI 27.007: "GSM w/EGPRS").
	MmModemAccessTechnologyUmts       MMModemAccessTechnology = 1 << 5     // UMTS (ETSI 27.007: "UTRAN").
	MmModemAccessTechnologyHsdpa      MMModemAccessTechnology = 1 << 6     // HSDPA (ETSI 27.007: "UTRAN w/HSDPA").
	MmModemAccessTechnologyHsupa      MMModemAccessTechnology = 1 << 7     // HSUPA (ETSI 27.007: "UTRAN w/HSUPA").
	MmModemAccessTechnologyHspa       MMModemAccessTechnology = 1 << 8     // HSPA (ETSI 27.007: "UTRAN w/HSDPA and HSUPA").
	MmModemAccessTechnologyHspaPlus   MMModemAccessTechnology = 1 << 9     // HSPA+ (ETSI 27.007: "UTRAN w/HSPA+").
	MmModemAccessTechnology1xrtt      MMModemAccessTechnology = 1 << 10    // CDMA2000 1xRTT.
	MmModemAccessTechnologyEvdo0      MMModemAccessTechnology = 1 << 11    // CDMA2000 EVDO revision 0.
	MmModemAccessTechnologyEvdoa      MMModemAccessTechnology = 1 << 12    // CDMA2000 EVDO revision A.
	MmModemAccessTechnologyEvdob      MMModemAccessTechnology = 1 << 13    // CDMA2000 EVDO revision B.
	MmModemAccessTechnologyLte        MMModemAccessTechnology = 1 << 14    // LTE (ETSI 27.007: "E-UTRAN")
	MmModemAccessTechnologyAny        MMModemAccessTechnology = 0xFFFFFFFF // Mask specifying all access technologies.
)

type MMModemMode uint32 // Bitfield to indicate which access modes are supported, allowed or preferred in a given device.

//go:generate stringer -type=MMModemMode
const (
	MmModemModeNone MMModemMode = 0          // None
	MmModemModeCs   MMModemMode = 1 << 0     // CSD, GSM, and other circuit-switched technologies.
	MmModemMode2g   MMModemMode = 1 << 1     // GPRS, EDGE.
	MmModemMode3g   MMModemMode = 1 << 2     // UMTS, HSxPA.
	MmModemMode4g   MMModemMode = 1 << 3     // LTE.
	MmModemModeAny  MMModemMode = 0xFFFFFFFF // Any mode can be used (only this value allowed for POTS modems).

)

type MMModemBand uint32 // Radio bands supported by the device when connecting to a mobile network.
//go:generate stringer -type=MMModemBand
const (
	MmModemBandUnknown MMModemBand = 0 // Unknown or invalid band.
	/* GSM/UMTS bands */
	MmModemBandEgsm   MMModemBand = 1  // GSM/GPRS/EDGE 900 MHz.
	MmModemBandDcs    MMModemBand = 2  // GSM/GPRS/EDGE 1800 MHz.
	MmModemBandPcs    MMModemBand = 3  // GSM/GPRS/EDGE 1900 MHz.
	MmModemBandG850   MMModemBand = 4  // GSM/GPRS/EDGE 850 MHz.
	MmModemBandUtran1 MMModemBand = 5  // UMTS 2100 MHz (IMT, UTRAN band 1).
	MmModemBandUtran3 MMModemBand = 6  // UMTS 1800 MHz (DCS, UTRAN band 3).
	MmModemBandUtran4 MMModemBand = 7  // UMTS 1700 MHz (AWS A-F, UTRAN band 4).
	MmModemBandUtran6 MMModemBand = 8  // UMTS 800 MHz (UTRAN band 6).
	MmModemBandUtran5 MMModemBand = 9  // UMTS 850 MHz (CLR, UTRAN band 5).
	MmModemBandUtran8 MMModemBand = 10 //UMTS 900 MHz (E-GSM, UTRAN band 8).
	MmModemBandUtran9 MMModemBand = 11 // UMTS 1700 MHz (UTRAN band 9).
	MmModemBandUtran2 MMModemBand = 12 //  UMTS 1900 MHz (PCS A-F, UTRAN band 2).
	MmModemBandUtran7 MMModemBand = 13 // UMTS 2600 MHz (IMT-E, UTRAN band 7).
	MmModemBandG450   MMModemBand = 14 // GSM/GPRS/EDGE 450 MHz.
	MmModemBandG480   MMModemBand = 15 // GSM/GPRS/EDGE 480 MHz.
	MmModemBandG750   MMModemBand = 16 // GSM/GPRS/EDGE 750 MHz.
	MmModemBandG380   MMModemBand = 17 // GSM/GPRS/EDGE 380 MHz.
	MmModemBandG410   MMModemBand = 18 // GSM/GPRS/EDGE 410 MHz.
	MmModemBandG710   MMModemBand = 19 // GSM/GPRS/EDGE 710 MHz.
	MmModemBandG810   MMModemBand = 20 // GSM/GPRS/EDGE 810 MHz.
	/* LTE bands */
	MmModemBandEutran1  MMModemBand = 31  // E-UTRAN band 1.
	MmModemBandEutran2  MMModemBand = 32  // E-UTRAN band 2.
	MmModemBandEutran3  MMModemBand = 33  // E-UTRAN band 3.
	MmModemBandEutran4  MMModemBand = 34  // E-UTRAN band 4.
	MmModemBandEutran5  MMModemBand = 35  // E-UTRAN band 5.
	MmModemBandEutran6  MMModemBand = 36  // E-UTRAN band 6.
	MmModemBandEutran7  MMModemBand = 37  // E-UTRAN band 7.
	MmModemBandEutran8  MMModemBand = 38  // E-UTRAN band 8.
	MmModemBandEutran9  MMModemBand = 39  // E-UTRAN band 9.
	MmModemBandEutran10 MMModemBand = 40  // E-UTRAN band 10.
	MmModemBandEutran11 MMModemBand = 41  // E-UTRAN band 11.
	MmModemBandEutran12 MMModemBand = 42  // E-UTRAN band 12.
	MmModemBandEutran13 MMModemBand = 43  // E-UTRAN band 13.
	MmModemBandEutran14 MMModemBand = 44  // E-UTRAN band 14.
	MmModemBandEutran17 MMModemBand = 47  // E-UTRAN band 17.
	MmModemBandEutran18 MMModemBand = 48  // E-UTRAN band 18.
	MmModemBandEutran19 MMModemBand = 49  // E-UTRAN band 19.
	MmModemBandEutran20 MMModemBand = 50  // E-UTRAN band 20.
	MmModemBandEutran21 MMModemBand = 51  // E-UTRAN band 21.
	MmModemBandEutran22 MMModemBand = 52  // E-UTRAN band 22.
	MmModemBandEutran23 MMModemBand = 53  // E-UTRAN band 23.
	MmModemBandEutran24 MMModemBand = 54  // E-UTRAN band 24.
	MmModemBandEutran25 MMModemBand = 55  // E-UTRAN band 25.
	MmModemBandEutran26 MMModemBand = 56  // E-UTRAN band 26.
	MmModemBandEutran27 MMModemBand = 57  // E-UTRAN band 27.
	MmModemBandEutran28 MMModemBand = 58  // E-UTRAN band 28.
	MmModemBandEutran29 MMModemBand = 59  // E-UTRAN band 29.
	MmModemBandEutran30 MMModemBand = 60  // E-UTRAN band 30.
	MmModemBandEutran31 MMModemBand = 61  // E-UTRAN band 31.
	MmModemBandEutran32 MMModemBand = 62  // E-UTRAN band 32.
	MmModemBandEutran33 MMModemBand = 63  // E-UTRAN band 33.
	MmModemBandEutran34 MMModemBand = 64  // E-UTRAN band 34.
	MmModemBandEutran35 MMModemBand = 65  // E-UTRAN band 35.
	MmModemBandEutran36 MMModemBand = 66  // E-UTRAN band 36.
	MmModemBandEutran37 MMModemBand = 67  // E-UTRAN band 37.
	MmModemBandEutran38 MMModemBand = 68  // E-UTRAN band 38.
	MmModemBandEutran39 MMModemBand = 69  // E-UTRAN band 39.
	MmModemBandEutran40 MMModemBand = 70  // E-UTRAN band 40
	MmModemBandEutran41 MMModemBand = 71  // E-UTRAN band 41.
	MmModemBandEutran42 MMModemBand = 72  // E-UTRAN band 42.
	MmModemBandEutran43 MMModemBand = 73  // E-UTRAN band 43.
	MmModemBandEutran44 MMModemBand = 74  // E-UTRAN band 44.
	MmModemBandEutran45 MMModemBand = 75  // E-UTRAN band 45.
	MmModemBandEutran46 MMModemBand = 76  // E-UTRAN band 46.
	MmModemBandEutran47 MMModemBand = 77  // E-UTRAN band 47.
	MmModemBandEutran48 MMModemBand = 78  // E-UTRAN band 48.
	MmModemBandEutran49 MMModemBand = 79  // E-UTRAN band 49.
	MmModemBandEutran50 MMModemBand = 80  // E-UTRAN band 50.
	MmModemBandEutran51 MMModemBand = 81  // E-UTRAN band 51.
	MmModemBandEutran52 MMModemBand = 82  // E-UTRAN band 52.
	MmModemBandEutran53 MMModemBand = 83  // E-UTRAN band 53.
	MmModemBandEutran54 MMModemBand = 84  // E-UTRAN band 54.
	MmModemBandEutran55 MMModemBand = 85  // E-UTRAN band 55.
	MmModemBandEutran56 MMModemBand = 86  // E-UTRAN band 56.
	MmModemBandEutran57 MMModemBand = 87  // E-UTRAN band 57.
	MmModemBandEutran58 MMModemBand = 88  // E-UTRAN band 58.
	MmModemBandEutran59 MMModemBand = 89  // E-UTRAN band 59.
	MmModemBandEutran60 MMModemBand = 90  // E-UTRAN band 60.
	MmModemBandEutran61 MMModemBand = 91  // E-UTRAN band 61.
	MmModemBandEutran62 MMModemBand = 92  // E-UTRAN band 62.
	MmModemBandEutran63 MMModemBand = 93  // E-UTRAN band 63.
	MmModemBandEutran64 MMModemBand = 94  // E-UTRAN band 64.
	MmModemBandEutran65 MMModemBand = 95  // E-UTRAN band 65.
	MmModemBandEutran66 MMModemBand = 96  // E-UTRAN band 66.
	MmModemBandEutran67 MMModemBand = 97  // E-UTRAN band 67.
	MmModemBandEutran68 MMModemBand = 98  // E-UTRAN band 68.
	MmModemBandEutran69 MMModemBand = 99  // E-UTRAN band 69.
	MmModemBandEutran70 MMModemBand = 100 // E-UTRAN band 70.
	MmModemBandEutran71 MMModemBand = 101 // E-UTRAN band 71.
	/* CDMA Band Classes (see 3GPP2 C.S0057-C) */
	MmModemBandCdmaBc0  MMModemBand = 128 // CDMA Band Class 0 (US Cellular 850MHz).
	MmModemBandCdmaBc1  MMModemBand = 129 // CDMA Band Class 1 (US PCS 1900MHz).
	MmModemBandCdmaBc2  MMModemBand = 130 // CDMA Band Class 2 (UK TACS 900MHz).
	MmModemBandCdmaBc3  MMModemBand = 131 // CDMA Band Class 3 (Japanese TACS).
	MmModemBandCdmaBc4  MMModemBand = 132 // CDMA Band Class 4 (Korean PCS).
	MmModemBandCdmaBc5  MMModemBand = 134 // CDMA Band Class 5 (NMT 450MHz).
	MmModemBandCdmaBc6  MMModemBand = 135 // CDMA Band Class 6 (IMT2000 2100MHz).
	MmModemBandCdmaBc7  MMModemBand = 136 // CDMA Band Class 7 (Cellular 700MHz).
	MmModemBandCdmaBc8  MMModemBand = 137 // CDMA Band Class 8 (1800MHz).
	MmModemBandCdmaBc9  MMModemBand = 138 // CDMA Band Class 9 (900MHz).
	MmModemBandCdmaBc10 MMModemBand = 139 // (US Secondary 800).
	MmModemBandCdmaBc11 MMModemBand = 140 // (European PAMR 400MHz).
	MmModemBandCdmaBc12 MMModemBand = 141 // (PAMR 800MHz).
	MmModemBandCdmaBc13 MMModemBand = 142 // (IMT2000 2500MHz Expansion).
	MmModemBandCdmaBc14 MMModemBand = 143 // (More US PCS 1900MHz).
	MmModemBandCdmaBc15 MMModemBand = 144 // (AWS 1700MHz).
	MmModemBandCdmaBc16 MMModemBand = 145 // (US 2500MHz).
	MmModemBandCdmaBc17 MMModemBand = 146 // (US 2500MHz Forward Link Only).
	MmModemBandCdmaBc18 MMModemBand = 147 // (US 700MHz Public Safety).
	MmModemBandCdmaBc19 MMModemBand = 148 // (US Lower 700MHz).
	/* Additional UMTS bands:
	 *  15-18 reserved
	 *  23-24 reserved
	 *  27-31 reserved
	 */
	MmModemBandUtran10 MMModemBand = 210 // UMTS 1700 MHz (EAWS A-G, UTRAN band 10).
	MmModemBandUtran11 MMModemBand = 211 // UMTS 1500 MHz (LPDC, UTRAN band 11).
	MmModemBandUtran12 MMModemBand = 212 // UMTS 700 MHz (LSMH A/B/C, UTRAN band 12)
	MmModemBandUtran13 MMModemBand = 213 // UMTS 700 MHz (USMH C, UTRAN band 13)
	MmModemBandUtran14 MMModemBand = 214 // UMTS 700 MHz (USMH D, UTRAN band 14)
	MmModemBandUtran19 MMModemBand = 219 // UMTS 800 MHz (UTRAN band 19).
	MmModemBandUtran20 MMModemBand = 220 // UMTS 800 MHz (EUDD, UTRAN band 20).
	MmModemBandUtran21 MMModemBand = 221 // UMTS 1500 MHz (UPDC, UTRAN band 21).
	MmModemBandUtran22 MMModemBand = 222 // UMTS 3500 MHz (UTRAN band 22).
	MmModemBandUtran25 MMModemBand = 225 // UMTS 1900 MHz (EPCS A-G, UTRAN band 25).
	MmModemBandUtran26 MMModemBand = 226 // UMTS 850 MHz (ECLR, UTRAN band 26).
	MmModemBandUtran32 MMModemBand = 232 // UMTS 1500 MHz (L-band, UTRAN band 32).
	/* All/Any */
	MmModemBandAny MMModemBand = 256 // For certain operations, allow the modem to select a band automatically.
)

type MMModemPortType uint32 // Type of modem port.
//go:generate stringer -type=MMModemPortType
const (
	MmModemPortTypeUnknown MMModemPortType = 1 // Unknown.
	MmModemPortTypeNet     MMModemPortType = 2 // Net port.
	MmModemPortTypeAt      MMModemPortType = 3 // AT port.
	MmModemPortTypeQcdm    MMModemPortType = 4 // QCDM port.
	MmModemPortTypeGps     MMModemPortType = 5 // GPS port.
	MmModemPortTypeQmi     MMModemPortType = 6 // QMI port.
	MmModemPortTypeMbim    MMModemPortType = 7 // MBIM port.
	MmModemPortTypeAudio   MMModemPortType = 8 // Audio port.

)

type MMSmsPduType uint32 // Type of PDUs used in the SMS.
//go:generate stringer -type=MMSmsPduType
const (
	MmSmsPduTypeUnknown                     MMSmsPduType = 0  // Unknown type.
	MmSmsPduTypeDeliver                     MMSmsPduType = 1  // 3GPP Mobile-Terminated (MT) message.
	MmSmsPduTypeSubmit                      MMSmsPduType = 2  // 3GPP Mobile-Originated (MO) message.
	MmSmsPduTypeStatusReport                MMSmsPduType = 3  // 3GPP status report (MT).
	MmSmsPduTypeCdmaDeliver                 MMSmsPduType = 32 // 3GPP2 Mobile-Terminated (MT) message.
	MmSmsPduTypeCdmaSubmit                  MMSmsPduType = 33 // 3GPP2 Mobile-Originated (MO) message.
	MmSmsPduTypeCdmaCancellation            MMSmsPduType = 34 // 3GPP2 Cancellation (MO) message.
	MmSmsPduTypeCdmaDeliveryAcknowledgement MMSmsPduType = 35 // 3GPP2 Delivery Acknowledgement (MT) message.
	MmSmsPduTypeCdmaUserAcknowledgement     MMSmsPduType = 36 // 3GPP2 User Acknowledgement (MT or MO) message.
	MmSmsPduTypeCdmaReadAcknowledgement     MMSmsPduType = 37 // 3GPP2 Read Acknowledgement (MT or MO) message.

)

type MMSmsState uint32 // State of a given SMS.
//go:generate stringer -type=MMSmsState
const (
	MmSmsStateUnknown   MMSmsState = 0 // State unknown or not reportable.
	MmSmsStateStored    MMSmsState = 1 // The message has been neither received nor yet sent.
	MmSmsStateReceiving MMSmsState = 2 // The message is being received but is not yet complete.
	MmSmsStateReceived  MMSmsState = 3 // The message has been completely received.
	MmSmsStateSending   MMSmsState = 4 // The message is queued for delivery.
	MmSmsStateSent      MMSmsState = 5 // The message was successfully sent.

)

type MMSmsDeliveryState uint32 // Known SMS delivery states as defined in 3GPP TS 03.40 and  3GPP2 N.S0005-O, section 6.5.2.125. States out of the known ranges may also be valid (either reserved or SC-specific).

//go:generate stringer -type=MMSmsDeliveryState
const (
	/* Completed deliveries */
	MmSmsDeliveryStateCompletedReceived             MMSmsDeliveryState = 0x00 // Delivery completed, message received by the SME.
	MmSmsDeliveryStateCompletedForwardedUnconfirmed MMSmsDeliveryState = 0x01 // Forwarded by the SC to the SME but the SC is unable to confirm delivery.
	MmSmsDeliveryStateCompletedReplacedBySc         MMSmsDeliveryState = 0x02 // Message replaced by the SC.

	/* Temporary failures */
	MmSmsDeliveryStateTemporaryErrorCongestion        MMSmsDeliveryState = 0x20 // Temporary error, congestion.
	MmSmsDeliveryStateTemporaryErrorSmeBusy           MMSmsDeliveryState = 0x21 // Temporary error, SME busy.
	MmSmsDeliveryStateTemporaryErrorNoResponseFromSme MMSmsDeliveryState = 0x22 // Temporary error, no response from the SME.
	MmSmsDeliveryStateTemporaryErrorServiceRejected   MMSmsDeliveryState = 0x23 // Temporary error, service rejected.
	MmSmsDeliveryStateTemporaryErrorQosNotAvailable   MMSmsDeliveryState = 0x24 // Temporary error, QoS not available.
	MmSmsDeliveryStateTemporaryErrorInSme             MMSmsDeliveryState = 0x25 // Temporary error in the SME.

	/* Permanent failures */
	MmSmsDeliveryStateErrorRemoteProcedure           MMSmsDeliveryState = 0x40 // Permanent remote procedure error.
	MmSmsDeliveryStateErrorIncompatibleDestination   MMSmsDeliveryState = 0x41 // Permanent error, incompatible destination.
	MmSmsDeliveryStateErrorConnectionRejected        MMSmsDeliveryState = 0x42 // Permanent error, connection rejected by the SME.
	MmSmsDeliveryStateErrorNotObtainable             MMSmsDeliveryState = 0x43 // Permanent error, not obtainable.
	MmSmsDeliveryStateErrorQosNotAvailable           MMSmsDeliveryState = 0x44 // Permanent error, QoS not available.
	MmSmsDeliveryStateErrorNoInterworkingAvailable   MMSmsDeliveryState = 0x45 // Permanent error, no interworking available.
	MmSmsDeliveryStateErrorValidityPeriodExpired     MMSmsDeliveryState = 0x46 // Permanent error, message validity period expired.
	MmSmsDeliveryStateErrorDeletedByOriginatingSme   MMSmsDeliveryState = 0x47 // Permanent error, deleted by originating SME.
	MmSmsDeliveryStateErrorDeletedByScAdministration MMSmsDeliveryState = 0x48 // Permanent error, deleted by SC administration.
	MmSmsDeliveryStateErrorMessageDoesNotExist       MMSmsDeliveryState = 0x49 // Permanent error, message does no longer exist.

	/* Temporary failures that became permanent */
	MmSmsDeliveryStateTemporaryFatalErrorCongestion        MMSmsDeliveryState = 0x60 // Permanent error, congestion.
	MmSmsDeliveryStateTemporaryFatalErrorSmeBusy           MMSmsDeliveryState = 0x61 // Permanent error, SME busy.
	MmSmsDeliveryStateTemporaryFatalErrorNoResponseFromSme MMSmsDeliveryState = 0x62 // Permanent error, no response from the SME.
	MmSmsDeliveryStateTemporaryFatalErrorServiceRejected   MMSmsDeliveryState = 0x63 // Permanent error, service rejected.
	MmSmsDeliveryStateTemporaryFatalErrorQosNotAvailable   MMSmsDeliveryState = 0x64 // Permanent error, QoS not available.
	MmSmsDeliveryStateTemporaryFatalErrorInSme             MMSmsDeliveryState = 0x65 // Permanent error in SME.

	/* Unknown, out of any possible valid value [0x00-0xFF] */
	MmSmsDeliveryStateUnknown = 0x100 // Unknown state.

	/* --------------- 3GPP2 specific errors ---------------------- */

	/* Network problems */
	MmSmsDeliveryStateNetworkProblemAddressVacant             MMSmsDeliveryState = 0x200 // Permanent error in network, address vacant.
	MmSmsDeliveryStateNetworkProblemAddressTranslationFailure MMSmsDeliveryState = 0x201 // Permanent error in network, address translation failure.
	MmSmsDeliveryStateNetworkProblemNetworkResourceOutage     MMSmsDeliveryState = 0x202 // Permanent error in network, network resource outage.
	MmSmsDeliveryStateNetworkProblemNetworkFailure            MMSmsDeliveryState = 0x203 // Permanent error in network, network failure.
	MmSmsDeliveryStateNetworkProblemInvalidTeleserviceId      MMSmsDeliveryState = 0x204 // Permanent error in network, invalid teleservice id.
	MmSmsDeliveryStateNetworkProblemOther                     MMSmsDeliveryState = 0x205 // Permanent error, other network problem.
	/* Terminal problems */
	MmSmsDeliveryStateTerminalProblemNoPageResponse                   MMSmsDeliveryState = 0x220 // Permanent error in terminal, no page response.
	MmSmsDeliveryStateTerminalProblemDestinationBusy                  MMSmsDeliveryState = 0x221 // Permanent error in terminal, destination busy.
	MmSmsDeliveryStateTerminalProblemNoAcknowledgment                 MMSmsDeliveryState = 0x222 // Permanent error in terminal, no acknowledgement.
	MmSmsDeliveryStateTerminalProblemDestinationResourceShortage      MMSmsDeliveryState = 0x223 // Permanent error in terminal, destination resource shortage.
	MmSmsDeliveryStateTerminalProblemSmsDeliveryPostponed             MMSmsDeliveryState = 0x224 // Permanent error in terminal, SMS delivery postponed.
	MmSmsDeliveryStateTerminalProblemDestinationOutOfService          MMSmsDeliveryState = 0x225 // Permanent error in terminal, destination out of service.
	MmSmsDeliveryStateTerminalProblemDestinationNoLongerAtThisAddress MMSmsDeliveryState = 0x226 // Permanent error in terminal, destination no longer at this address.
	MmSmsDeliveryStateTerminalProblemOther                            MMSmsDeliveryState = 0x227 // Permanent error, other terminal problem.
	/* Radio problems */
	MmSmsDeliveryStateRadioInterfaceProblemResourceShortage MMSmsDeliveryState = 0x240 // Permanent error in radio interface, resource shortage.
	MmSmsDeliveryStateRadioInterfaceProblemIncompatibility  MMSmsDeliveryState = 0x241 // Permanent error in radio interface, problem incompatibility.
	MmSmsDeliveryStateRadioInterfaceProblemOther            MMSmsDeliveryState = 0x242 // Permanent error, other radio interface problem.
	/* General problems */
	MmSmsDeliveryStateGeneralProblemEncoding                         MMSmsDeliveryState = 0x260 // Permanent error, encoding.
	MmSmsDeliveryStateGeneralProblemSmsOriginationDenied             MMSmsDeliveryState = 0x261 // Permanent error, SMS origination denied.
	MmSmsDeliveryStateGeneralProblemSmsTerminationDenied             MMSmsDeliveryState = 0x262 // Permanent error, SMS termination denied.
	MmSmsDeliveryStateGeneralProblemSupplementaryServiceNotSupported MMSmsDeliveryState = 0x263 // Permanent error, supplementary service not supported.
	MmSmsDeliveryStateGeneralProblemSmsNotSupported                  MMSmsDeliveryState = 0x264 // Permanent error, SMS not supported.
	MmSmsDeliveryStateGeneralProblemMissingExpectedParameter         MMSmsDeliveryState = 0x266 // Permanent error, missing expected parameter.
	MmSmsDeliveryStateGeneralProblemMissingMandatoryParameter        MMSmsDeliveryState = 0x267 // Permanent error, missing mandatory parameter.
	MmSmsDeliveryStateGeneralProblemUnrecognizedParameterValue       MMSmsDeliveryState = 0x268 // Permanent error, unrecognized parameter value.
	MmSmsDeliveryStateGeneralProblemUnexpectedParameterValue         MMSmsDeliveryState = 0x269 // Permanent error, unexpected parameter value.
	MmSmsDeliveryStateGeneralProblemUserDataSizeError                MMSmsDeliveryState = 0x26A // Permanent error, user data size error.
	MmSmsDeliveryStateGeneralProblemOther                            MMSmsDeliveryState = 0x26B //  Permanent error, other general problem.

	/* Temporary network problems */
	MmSmsDeliveryStateTemporaryNetworkProblemAddressVacant             MMSmsDeliveryState = 0x300 // Temporary error in network, address vacant.
	MmSmsDeliveryStateTemporaryNetworkProblemAddressTranslationFailure MMSmsDeliveryState = 0x301 // Temporary error in network, address translation failure.
	MmSmsDeliveryStateTemporaryNetworkProblemNetworkResourceOutage     MMSmsDeliveryState = 0x302 // Temporary error in network, network resource outage.
	MmSmsDeliveryStateTemporaryNetworkProblemNetworkFailure            MMSmsDeliveryState = 0x303 // Temporary error in network, network failure.
	MmSmsDeliveryStateTemporaryNetworkProblemInvalidTeleserviceId      MMSmsDeliveryState = 0x304 // Temporary error in network, invalid teleservice id.
	MmSmsDeliveryStateTemporaryNetworkProblemOther                     MMSmsDeliveryState = 0x305 // Temporary error, other network problem.
	/* Temporary terminal problems */
	MmSmsDeliveryStateTemporaryTerminalProblemNoPageResponse                   MMSmsDeliveryState = 0x320 // Temporary error in terminal, no page response.
	MmSmsDeliveryStateTemporaryTerminalProblemDestinationBusy                  MMSmsDeliveryState = 0x321 // Temporary error in terminal, destination busy.
	MmSmsDeliveryStateTemporaryTerminalProblemNoAcknowledgment                 MMSmsDeliveryState = 0x322 // Temporary error in terminal, no acknowledgement.
	MmSmsDeliveryStateTemporaryTerminalProblemDestinationResourceShortage      MMSmsDeliveryState = 0x323 // Temporary error in terminal, destination resource shortage.
	MmSmsDeliveryStateTemporaryTerminalProblemSmsDeliveryPostponed             MMSmsDeliveryState = 0x324 // Temporary error in terminal, SMS delivery postponed.
	MmSmsDeliveryStateTemporaryTerminalProblemDestinationOutOfService          MMSmsDeliveryState = 0x325 // Temporary error in terminal, destination out of service.
	MmSmsDeliveryStateTemporaryTerminalProblemDestinationNoLongerAtThisAddress MMSmsDeliveryState = 0x326 // Temporary error in terminal, destination no longer at this address.
	MmSmsDeliveryStateTemporaryTerminalProblemOther                            MMSmsDeliveryState = 0x327 // Temporary error, other terminal problem.
	/* Temporary radio problems */
	MmSmsDeliveryStateTemporaryRadioInterfaceProblemResourceShortage MMSmsDeliveryState = 0x340 // Temporary error in radio interface, resource shortage.
	MmSmsDeliveryStateTemporaryRadioInterfaceProblemIncompatibility  MMSmsDeliveryState = 0x341 // Temporary error in radio interface, problem incompatibility.
	MmSmsDeliveryStateTemporaryRadioInterfaceProblemOther            MMSmsDeliveryState = 0x342 // Temporary error, other radio interface problem.
	/* Temporary general problems */
	MmSmsDeliveryStateTemporaryGeneralProblemEncoding                         MMSmsDeliveryState = 0x360 // Temporary error, encoding.
	MmSmsDeliveryStateTemporaryGeneralProblemSmsOriginationDenied             MMSmsDeliveryState = 0x361 // Temporary error, SMS origination denied.
	MmSmsDeliveryStateTemporaryGeneralProblemSmsTerminationDenied             MMSmsDeliveryState = 0x362 // Temporary error, SMS termination denied.
	MmSmsDeliveryStateTemporaryGeneralProblemSupplementaryServiceNotSupported MMSmsDeliveryState = 0x363 // Temporary error, supplementary service not supported.
	MmSmsDeliveryStateTemporaryGeneralProblemSmsNotSupported                  MMSmsDeliveryState = 0x364 // Temporary error, SMS not supported.
	MmSmsDeliveryStateTemporaryGeneralProblemMissingExpectedParameter         MMSmsDeliveryState = 0x366 // Temporary error, missing expected parameter.
	MmSmsDeliveryStateTemporaryGeneralProblemMissingMandatoryParameter        MMSmsDeliveryState = 0x367 // Temporary error, missing mandatory parameter.
	MmSmsDeliveryStateTemporaryGeneralProblemUnrecognizedParameterValue       MMSmsDeliveryState = 0x368 // Temporary error, unrecognized parameter value.
	MmSmsDeliveryStateTemporaryGeneralProblemUnexpectedParameterValue         MMSmsDeliveryState = 0x369 // Temporary error, unexpected parameter value.
	MmSmsDeliveryStateTemporaryGeneralProblemUserDataSizeError                MMSmsDeliveryState = 0x36A // Temporary error, user data size error.
	MmSmsDeliveryStateTemporaryGeneralProblemOther                            MMSmsDeliveryState = 0x36B // Temporary error, other general problem.

)

type MMSmsStorage uint32 // Storage for SMS messages.

//go:generate stringer -type=MMSmsStorage
const (
	MmSmsStorageUnknown MMSmsStorage = 0 // Storage unknown.
	MmSmsStorageSm      MMSmsStorage = 1 // SIM card storage area.
	MmSmsStorageMe      MMSmsStorage = 2 // Mobile equipment storage area.
	MmSmsStorageMt      MMSmsStorage = 3 // Sum of SIM and Mobile equipment storages
	MmSmsStorageSr      MMSmsStorage = 4 // Status report message storage area.
	MmSmsStorageBm      MMSmsStorage = 5 // Broadcast message storage area.
	MmSmsStorageTa      MMSmsStorage = 6 // Terminal adaptor message storage area.

)

type MMSmsValidityType uint32 // Type of SMS validity value.

//go:generate stringer -type=MMSmsValidityType
const (
	MmSmsValidityTypeUnknown  MMSmsValidityType = 0 // Validity type unknown.
	MmSmsValidityTypeRelative MMSmsValidityType = 1 // Relative validity.
	MmSmsValidityTypeAbsolute MMSmsValidityType = 2 // Absolute validity.
	MmSmsValidityTypeEnhanced MMSmsValidityType = 3 // Enhanced validity.

)

type MMSmsCdmaTeleserviceId uint32 // Teleservice IDs supported for CDMA SMS, as defined in 3GPP2 X.S0004-550-E (section 2.256) and 3GPP2 C.S0015-B (section 3.4.3.1).

//go:generate stringer -type=MMSmsCdmaTeleserviceId
const (
	MmSmsCdmaTeleserviceIdUnknown MMSmsCdmaTeleserviceId = 0x0000 // Unknown.
	MmSmsCdmaTeleserviceIdCmt91   MMSmsCdmaTeleserviceId = 0x1000 // IS-91 Extended Protocol Enhanced Services.
	MmSmsCdmaTeleserviceIdWpt     MMSmsCdmaTeleserviceId = 0x1001 // Wireless Paging Teleservice.
	MmSmsCdmaTeleserviceIdWmt     MMSmsCdmaTeleserviceId = 0x1002 // Wireless Messaging Teleservice.
	MmSmsCdmaTeleserviceIdVmn     MMSmsCdmaTeleserviceId = 0x1003 // Voice Mail Notification.
	MmSmsCdmaTeleserviceIdWap     MMSmsCdmaTeleserviceId = 0x1004 // Wireless Application Protocol.
	MmSmsCdmaTeleserviceIdWemt    MMSmsCdmaTeleserviceId = 0x1005 // Wireless Enhanced Messaging Teleservice.
	MmSmsCdmaTeleserviceIdScpt    MMSmsCdmaTeleserviceId = 0x1006 // Service Category Programming Teleservice
	MmSmsCdmaTeleserviceIdCatpt   MMSmsCdmaTeleserviceId = 0x1007 // Card Application Toolkit Protocol Teleservice.

)

type MMSmsCdmaServiceCategory uint32 // Service category for CDMA SMS, as defined in 3GPP2 C.R1001-D (section 9.3).

//go:generate stringer -type=MMSmsCdmaServiceCategory
const (
	MmSmsCdmaServiceCategoryUnknown                        MMSmsCdmaServiceCategory = 0x0000 // Unknown.
	MmSmsCdmaServiceCategoryEmergencyBroadcast             MMSmsCdmaServiceCategory = 0x0001 // Emergency broadcast.
	MmSmsCdmaServiceCategoryAdministrative                 MMSmsCdmaServiceCategory = 0x0002 // Administrative.
	MmSmsCdmaServiceCategoryMaintenance                    MMSmsCdmaServiceCategory = 0x0003 // Maintenance.
	MmSmsCdmaServiceCategoryGeneralNewsLocal               MMSmsCdmaServiceCategory = 0x0004 // General news (local).
	MmSmsCdmaServiceCategoryGeneralNewsRegional            MMSmsCdmaServiceCategory = 0x0005 // General news (regional).
	MmSmsCdmaServiceCategoryGeneralNewsNational            MMSmsCdmaServiceCategory = 0x0006 // General news (national).
	MmSmsCdmaServiceCategoryGeneralNewsInternational       MMSmsCdmaServiceCategory = 0x0007 // General news (international).
	MmSmsCdmaServiceCategoryBusinessNewsLocal              MMSmsCdmaServiceCategory = 0x0008 // Business/Financial news (local).
	MmSmsCdmaServiceCategoryBusinessNewsRegional           MMSmsCdmaServiceCategory = 0x0009 // Business/Financial news (regional).
	MmSmsCdmaServiceCategoryBusinessNewsNational           MMSmsCdmaServiceCategory = 0x000A // Business/Financial news (national).
	MmSmsCdmaServiceCategoryBusinessNewsInternational      MMSmsCdmaServiceCategory = 0x000B // Business/Financial news (international).
	MmSmsCdmaServiceCategorySportsNewsLocal                MMSmsCdmaServiceCategory = 0x000C // Sports news (local).
	MmSmsCdmaServiceCategorySportsNewsRegional             MMSmsCdmaServiceCategory = 0x000D // Sports news (regional).
	MmSmsCdmaServiceCategorySportsNewsNational             MMSmsCdmaServiceCategory = 0x000E // Sports news (national).
	MmSmsCdmaServiceCategorySportsNewsInternational        MMSmsCdmaServiceCategory = 0x000F // Sports news (international).
	MmSmsCdmaServiceCategoryEntertainmentNewsLocal         MMSmsCdmaServiceCategory = 0x0010 // Entertainment news (local).
	MmSmsCdmaServiceCategoryEntertainmentNewsRegional      MMSmsCdmaServiceCategory = 0x0011 // Entertainment news (regional).
	MmSmsCdmaServiceCategoryEntertainmentNewsNational      MMSmsCdmaServiceCategory = 0x0012 // Entertainment news (national).
	MmSmsCdmaServiceCategoryEntertainmentNewsInternational MMSmsCdmaServiceCategory = 0x0013 // Entertainment news (international).
	MmSmsCdmaServiceCategoryLocalWeather                   MMSmsCdmaServiceCategory = 0x0014 // Local weather.
	MmSmsCdmaServiceCategoryTrafficReport                  MMSmsCdmaServiceCategory = 0x0015 // Area traffic report.
	MmSmsCdmaServiceCategoryFlightSchedules                MMSmsCdmaServiceCategory = 0x0016 // Local airport flight schedules.
	MmSmsCdmaServiceCategoryRestaurants                    MMSmsCdmaServiceCategory = 0x0017 // Restaurants.
	MmSmsCdmaServiceCategoryLodgings                       MMSmsCdmaServiceCategory = 0x0018 // Lodgings.
	MmSmsCdmaServiceCategoryRetailDirectory                MMSmsCdmaServiceCategory = 0x0019 // Retail directory.
	MmSmsCdmaServiceCategoryAdvertisements                 MMSmsCdmaServiceCategory = 0x001A // Advertisements.
	MmSmsCdmaServiceCategoryStockQuotes                    MMSmsCdmaServiceCategory = 0x001B // Stock quotes.
	MmSmsCdmaServiceCategoryEmployment                     MMSmsCdmaServiceCategory = 0x001C // Employment.
	MmSmsCdmaServiceCategoryHospitals                      MMSmsCdmaServiceCategory = 0x001D // Medical / Health / Hospitals.
	MmSmsCdmaServiceCategoryTechnologyNews                 MMSmsCdmaServiceCategory = 0x001E // Technology news.
	MmSmsCdmaServiceCategoryMulticategory                  MMSmsCdmaServiceCategory = 0x001F // Multi-category.
	MmSmsCdmaServiceCategoryCmasPresidentialAlert          MMSmsCdmaServiceCategory = 0x1000 // Presidential alert.
	MmSmsCdmaServiceCategoryCmasExtremeThreat              MMSmsCdmaServiceCategory = 0x1001 // Extreme threat.
	MmSmsCdmaServiceCategoryCmasSevereThreat               MMSmsCdmaServiceCategory = 0x1002 // Severe threat.
	MmSmsCdmaServiceCategoryCmasChildAbductionEmergency    MMSmsCdmaServiceCategory = 0x1003 // Child abduction emergency.
	MmSmsCdmaServiceCategoryCmasTest                       MMSmsCdmaServiceCategory = 0x1004 // CMAS test.

)

type MMModemLocationSource uint32 // Sources of location information supported by the modem.

//go:generate stringer -type=MMModemLocationSource
const (
	MmModemLocationSourceNone         MMModemLocationSource = 0      // None.
	MmModemLocationSource3gppLacCi    MMModemLocationSource = 1 << 0 //  Location Area Code and Cell ID.
	MmModemLocationSourceGpsRaw       MMModemLocationSource = 1 << 1 // GPS location given by predefined keys.
	MmModemLocationSourceGpsNmea      MMModemLocationSource = 1 << 2 // GPS location given as NMEA traces.
	MmModemLocationSourceCdmaBs       MMModemLocationSource = 1 << 3 // CDMA base station position.
	MmModemLocationSourceGpsUnmanaged MMModemLocationSource = 1 << 4 // No location given, just GPS module setup.
	MmModemLocationSourceAgpsMsa      MMModemLocationSource = 1 << 5 // Mobile Station Assisted A-GPS location requested.
	MmModemLocationSourceAgpsMsb      MMModemLocationSource = 1 << 6 // Mobile Station Based A-GPS location requested.

)

type MMModemLocationAssistanceDataType uint32 // Type of assistance data that may be injected to the GNSS module.

//go:generate stringer -type=MMModemLocationAssistanceDataType
const (
	MmModemLocationAssistanceDataTypeNone MMModemLocationAssistanceDataType = 0      // None.
	MmModemLocationAssistanceDataTypeXtra MMModemLocationAssistanceDataType = 1 << 0 // Qualcomm gpsOneXTRA.
)

type MMModemContactsStorage uint32 // Specifies different storage locations for contact information.

//go:generate stringer -type=MMModemContactsStorage
const (
	MmModemContactsStorageUnknown MMModemContactsStorage = 0 // Unknown location.
	MmModemContactsStorageMe      MMModemContactsStorage = 1 // Device's local memory.
	MmModemContactsStorageSm      MMModemContactsStorage = 2 // Card inserted in the device (like a SIM/RUIM).
	MmModemContactsStorageMt      MMModemContactsStorage = 3 // Combined device/ME and SIM/SM phonebook.

)

type MMBearerType uint32 // Type of context (2G/3G) or bearer (4G).

//go:generate stringer -type=MMBearerType
const (
	MmBearerTypeUnknown       MMBearerType = 0 // Unknown bearer.
	MmBearerTypeDefault       MMBearerType = 1 // Primary context (2G/3G) or default bearer (4G), defined by the user of the API.
	MmBearerTypeDefaultAttach MMBearerType = 2 // The initial default bearer established during LTE attach procedure, automatically connected as long as the device is  registered in the LTE network.
	MmBearerTypeDedicated     MMBearerType = 3 // Secondary context (2G/3G) or dedicated bearer (4G), defined by the user of the API. These bearers use the same IP address  used by a primary context or default bearer and provide a dedicated flow for  specific traffic with different QoS settings.
)

type MMBearerIpMethod uint32 // Type of IP method configuration to be used in a given Bearer.

//go:generate stringer -type=MMBearerIpMethod
const (
	MmBearerIpMethodUnknown MMBearerIpMethod = 0 // Unknown method.
	MmBearerIpMethodPpp     MMBearerIpMethod = 1 //  Use PPP to get IP addresses and DNS information. For IPv6, use PPP to retrieve the 64-bit Interface Identifier, use the IID to construct an IPv6 link-local address by following RFC 5072, and then run DHCP over the PPP link to retrieve DNS settings.
	MmBearerIpMethodStatic  MMBearerIpMethod = 2 // Use the provided static IP configuration given by the modem to configure the IP data interface.  Note that DNS servers may not be provided by the network or modem firmware.
	MmBearerIpMethodDhcp    MMBearerIpMethod = 3 // Begin DHCP or IPv6 SLAAC on the data interface to obtain any necessary IP configuration details that are not already provided by the IP configuration.  For IPv4 bearers DHCP should be used.  For IPv6 bearers SLAAC should be used, and the IP configuration may already contain a link-local address that should be assigned to the interface before SLAAC is started to obtain the rest of the configuration.

)

type MMBearerIpFamily uint32 // Type of IP family to be used in a given Bearer.

//go:generate stringer -type=MMBearerIpFamily
const (
	MmBearerIpFamilyNone   MMBearerIpFamily = 0          // None or unknown.
	MmBearerIpFamilyIpv4   MMBearerIpFamily = 1 << 0     // IPv4.
	MmBearerIpFamilyIpv6   MMBearerIpFamily = 1 << 1     // IPv6.
	MmBearerIpFamilyIpv4v6 MMBearerIpFamily = 1 << 2     // IPv4 and IPv6.
	MmBearerIpFamilyAny    MMBearerIpFamily = 0xFFFFFFFF // Mask specifying all IP families.

)

type MMBearerAllowedAuth uint32 // Allowed authentication methods when authenticating with the network.

//go:generate stringer -type=MMBearerAllowedAuth
const (
	MmBearerAllowedAuthUnknown MMBearerAllowedAuth = 0 // Unknown.
	/* bits 0..4 order match Ericsson device bitmap */
	MmBearerAllowedAuthNone     MMBearerAllowedAuth = 1 << 0 // None.
	MmBearerAllowedAuthPap      MMBearerAllowedAuth = 1 << 1 // PAP.
	MmBearerAllowedAuthChap     MMBearerAllowedAuth = 1 << 2 // CHAP.
	MmBearerAllowedAuthMschap   MMBearerAllowedAuth = 1 << 3 // MS-CHAP.
	MmBearerAllowedAuthMschapv2 MMBearerAllowedAuth = 1 << 4 // MS-CHAP v2.
	MmBearerAllowedAuthEap      MMBearerAllowedAuth = 1 << 5 // EAP.

)

type MMModemCdmaRegistrationState uint32 // Registration state of a CDMA modem.

//go:generate stringer -type=MMModemCdmaRegistrationState
const (
	MmModemCdmaRegistrationStateUnknown    MMModemCdmaRegistrationState = 0 // Registration status is unknown or the device is not registered.
	MmModemCdmaRegistrationStateRegistered MMModemCdmaRegistrationState = 1 // Registered, but roaming status is unknown or cannot be provided by the device. The device may or may not be roaming.
	MmModemCdmaRegistrationStateHome       MMModemCdmaRegistrationState = 2 // Currently registered on the home network.
	MmModemCdmaRegistrationStateRoaming    MMModemCdmaRegistrationState = 3 // Currently registered on a roaming network.

)

type MMModemCdmaActivationState uint32 // Activation state of a CDMA modem.

//go:generate stringer -type=MMModemCdmaActivationState
const (
	MmModemCdmaActivationStateUnknown            MMModemCdmaActivationState = 0 // Unknown activation state.
	MmModemCdmaActivationStateNotActivated       MMModemCdmaActivationState = 1 // Device is not activated
	MmModemCdmaActivationStateActivating         MMModemCdmaActivationState = 2 // Device is activating
	MmModemCdmaActivationStatePartiallyActivated MMModemCdmaActivationState = 3 // Device is partially activated; carrier-specific steps required to continue.
	MmModemCdmaActivationStateActivated          MMModemCdmaActivationState = 4 // Device is ready for use.

)

type MMModemCdmaRmProtocol uint32 // Protocol of the Rm interface in modems with CDMA capabilities.

//go:generate stringer -type=MMModemCdmaRmProtocol
const (
	MmModemCdmaRmProtocolUnknown           MMModemCdmaRmProtocol = 0 // Unknown protocol.
	MmModemCdmaRmProtocolAsync             MMModemCdmaRmProtocol = 1 // Asynchronous data or fax.
	MmModemCdmaRmProtocolPacketRelay       MMModemCdmaRmProtocol = 2 // Packet data service, Relay Layer Rm interface.
	MmModemCdmaRmProtocolPacketNetworkPpp  MMModemCdmaRmProtocol = 3 // Packet data service, Network Layer Rm interface, PPP.
	MmModemCdmaRmProtocolPacketNetworkSlip MMModemCdmaRmProtocol = 4 // Packet data service, Network Layer Rm interface, SLIP.
	MmModemCdmaRmProtocolStuIii            MMModemCdmaRmProtocol = 5 // STU-III service.

)

type MMModem3gppRegistrationState uint32 // GSM registration code as defined in 3GPP TS 27.007.

//go:generate stringer -type=MMModem3gppRegistrationState
const (
	MmModem3gppRegistrationStateIdle                    MMModem3gppRegistrationState = 0  // Not registered, not searching for new operator to register.
	MmModem3gppRegistrationStateHome                    MMModem3gppRegistrationState = 1  // Registered on home network.
	MmModem3gppRegistrationStateSearching               MMModem3gppRegistrationState = 2  // Not registered, searching for new operator to register with.
	MmModem3gppRegistrationStateDenied                  MMModem3gppRegistrationState = 3  // Registration denied.
	MmModem3gppRegistrationStateUnknown                 MMModem3gppRegistrationState = 4  // Unknown registration status.
	MmModem3gppRegistrationStateRoaming                 MMModem3gppRegistrationState = 5  // Registered on a roaming network.
	MmModem3gppRegistrationStateHomeSmsOnly             MMModem3gppRegistrationState = 6  // Registered for "SMS only", home network (applicable only when on LTE).
	MmModem3gppRegistrationStateRoamingSmsOnly          MMModem3gppRegistrationState = 7  // Registered for "SMS only", roaming network (applicable only when on LTE).
	MmModem3gppRegistrationStateEmergencyOnly           MMModem3gppRegistrationState = 8  // Emergency services only.
	MmModem3gppRegistrationStateHomeCsfbNotPreferred    MMModem3gppRegistrationState = 9  // Registered for "CSFB not preferred", home network (applicable only when on LTE).
	MmModem3gppRegistrationStateRoamingCsfbNotPreferred MMModem3gppRegistrationState = 10 // Registered for "CSFB not preferred", roaming network (applicable only when on LTE).

)

type MMModem3gppFacility uint32 // A bitfield describing which facilities have a lock enabled, i.e., requires a pin or unlock code. The facilities include the personalizations (device locks) described in 3GPP spec TS 22.022, and the PIN and PIN2 locks, which are SIM locks.

//go:generate stringer -type=MMModem3gppFacility
const (
	MmModem3gppFacilityNone         MMModem3gppFacility = 0      // No facility.
	MmModem3gppFacilitySim          MMModem3gppFacility = 1 << 0 // SIM lock.
	MmModem3gppFacilityFixedDialing MMModem3gppFacility = 1 << 1 // Fixed dialing (PIN2) SIM lock.
	MmModem3gppFacilityPhSim        MMModem3gppFacility = 1 << 2 // Device is locked to a specific SIM.
	MmModem3gppFacilityPhFsim       MMModem3gppFacility = 1 << 3 // Device is locked to first SIM inserted.
	MmModem3gppFacilityNetPers      MMModem3gppFacility = 1 << 4 // Network personalization.
	MmModem3gppFacilityNetSubPers   MMModem3gppFacility = 1 << 5 // Network subset personalization.
	MmModem3gppFacilityProviderPers MMModem3gppFacility = 1 << 6 // Service provider personalization.
	MmModem3gppFacilityCorpPers     MMModem3gppFacility = 1 << 7 // Corporate personalization.

)

type MMModem3gppNetworkAvailability uint32 // Network availability status as defined in 3GPP TS 27.007 section 7.3.

//go:generate stringer -type=MMModem3gppNetworkAvailability
const (
	MmModem3gppNetworkAvailabilityUnknown   MMModem3gppNetworkAvailability = 0 // Unknown availability.
	MmModem3gppNetworkAvailabilityAvailable MMModem3gppNetworkAvailability = 1 // Network is available.
	MmModem3gppNetworkAvailabilityCurrent   MMModem3gppNetworkAvailability = 2 // Network is the current one.
	MmModem3gppNetworkAvailabilityForbidden MMModem3gppNetworkAvailability = 3 // Network is forbidden.

)

type MMModem3gppSubscriptionState uint32 // Describes the current subscription status of the SIM.  This value is only available after the modem attempts to register with the network.

//go:generate stringer -type=MMModem3gppSubscriptionState
const (
	MmModem3gppSubscriptionStateUnknown       MMModem3gppSubscriptionState = 0 // The subscription state is unknown.
	MmModem3gppSubscriptionStateUnprovisioned MMModem3gppSubscriptionState = 1 // The account is unprovisioned.
	MmModem3gppSubscriptionStateProvisioned   MMModem3gppSubscriptionState = 2 // The account is provisioned and has data available.
	MmModem3gppSubscriptionStateOutOfData     MMModem3gppSubscriptionState = 3 // The account is provisioned but there is no data left.

)

type MMModem3gppUssdSessionState uint32 // State of a USSD session.

//go:generate stringer -type=MMModem3gppUssdSessionState
const (
	MmModem3gppUssdSessionStateUnknown      MMModem3gppUssdSessionState = 0 // Unknown state.
	MmModem3gppUssdSessionStateIdle         MMModem3gppUssdSessionState = 1 // No active session.
	MmModem3gppUssdSessionStateActive       MMModem3gppUssdSessionState = 2 // A session is active and the mobile is waiting for a response.
	MmModem3gppUssdSessionStateUserResponse MMModem3gppUssdSessionState = 3 // The network is waiting for the client's response.

)

type MMModem3gppEpsUeModeOperation uint32 // UE mode of operation for EPS, as per 3GPP TS 24.301.

//go:generate stringer -type=MMModem3gppEpsUeModeOperation
const (
	MmModem3gppEpsUeModeOperationUnknown MMModem3gppEpsUeModeOperation = 0 //  Unknown or not applicable.
	MmModem3gppEpsUeModeOperationPs1     MMModem3gppEpsUeModeOperation = 1 // PS mode 1 of operation: EPS only, voice-centric.
	MmModem3gppEpsUeModeOperationPs2     MMModem3gppEpsUeModeOperation = 2 // PS mode 2 of operation: EPS only, data-centric.
	MmModem3gppEpsUeModeOperationCsps1   MMModem3gppEpsUeModeOperation = 3 // CS/PS mode 1 of operation: EPS and non-EPS, voice-centric.
	MmModem3gppEpsUeModeOperationCsps2   MMModem3gppEpsUeModeOperation = 4 // CS/PS mode 2 of operation: EPS and non-EPS, data-centric.

)

type MMFirmwareImageType uint32 // Type of firmware image.

//go:generate stringer -type=MMFirmwareImageType
const (
	MmFirmwareImageTypeUnknown MMFirmwareImageType = 0 // Unknown firmware type.
	MmFirmwareImageTypeGeneric MMFirmwareImageType = 1 // Generic firmware image.
	MmFirmwareImageTypeGobi    MMFirmwareImageType = 2 // Firmware image of Gobi devices.

)

type MMOmaFeature uint32 // Features that can be enabled or disabled in the OMA device management support.

//go:generate stringer -type=MMOmaFeature
const (
	MmOmaFeatureNone                MMOmaFeature = 0      // None.
	MmOmaFeatureDeviceProvisioning  MMOmaFeature = 1 << 0 // Device provisioning service.
	MmOmaFeaturePrlUpdate           MMOmaFeature = 1 << 1 // PRL update service.
	MmOmaFeatureHandsFreeActivation MMOmaFeature = 1 << 2 // Hands free activation service.

)

type MMOmaSessionType uint32 // Type of OMA device management session.

//go:generate stringer -type=MMOmaSessionType
const (
	MmOmaSessionTypeUnknown                            MMOmaSessionType = 0  // Unknown session type.
	MmOmaSessionTypeClientInitiatedDeviceConfigure     MMOmaSessionType = 10 // Client-initiated device configure.
	MmOmaSessionTypeClientInitiatedPrlUpdate           MMOmaSessionType = 11 // Client-initiated PRL update.
	MmOmaSessionTypeClientInitiatedHandsFreeActivation MMOmaSessionType = 12 // Client-initiated hands free activation.
	MmOmaSessionTypeNetworkInitiatedDeviceConfigure    MMOmaSessionType = 20 // Network-initiated device configure.
	MmOmaSessionTypeNetworkInitiatedPrlUpdate          MMOmaSessionType = 21 // Network-initiated PRL update.
	MmOmaSessionTypeDeviceInitiatedPrlUpdate           MMOmaSessionType = 30 // Device-initiated PRL update.
	MmOmaSessionTypeDeviceInitiatedHandsFreeActivation MMOmaSessionType = 31 // Device-initiated hands free activation.

)

type MMOmaSessionState int32 // State of the OMA device management session.

//go:generate stringer -type=MMOmaSessionState
const (
	MmOmaSessionStateFailed               MMOmaSessionState = -1 // Failed.
	MmOmaSessionStateUnknown              MMOmaSessionState = 0  // Unknown.
	MmOmaSessionStateStarted              MMOmaSessionState = 1  // Started.
	MmOmaSessionStateRetrying             MMOmaSessionState = 2  // Retrying.
	MmOmaSessionStateConnecting           MMOmaSessionState = 3  // Connecting.
	MmOmaSessionStateConnected            MMOmaSessionState = 4  // Connected.
	MmOmaSessionStateAuthenticated        MMOmaSessionState = 5  // Authenticated.
	MmOmaSessionStateMdnDownloaded        MMOmaSessionState = 10 // MDN downloaded.
	MmOmaSessionStateMsidDownloaded       MMOmaSessionState = 11 // MSID downloaded.
	MmOmaSessionStatePrlDownloaded        MMOmaSessionState = 12 // PRL downloaded.
	MmOmaSessionStateMipProfileDownloaded MMOmaSessionState = 13 // MIP profile downloaded.
	MmOmaSessionStateCompleted            MMOmaSessionState = 20 // Session completed.

)

type MMOmaSessionStateFailedReason uint32 // Reason of failure in the OMA device management session.

//go:generate stringer -type=MMOmaSessionStateFailedReason
const (
	MmOmaSessionStateFailedReasonUnknown              MMOmaSessionStateFailedReason = 0 // No reason or unknown.
	MmOmaSessionStateFailedReasonNetworkUnavailable   MMOmaSessionStateFailedReason = 1 // Network unavailable.
	MmOmaSessionStateFailedReasonServerUnavailable    MMOmaSessionStateFailedReason = 2 // Server unavailable.
	MmOmaSessionStateFailedReasonAuthenticationFailed MMOmaSessionStateFailedReason = 3 // Authentication failed.
	MmOmaSessionStateFailedReasonMaxRetryExceeded     MMOmaSessionStateFailedReason = 4 // Maximum retries exceeded.
	MmOmaSessionStateFailedReasonSessionCancelled     MMOmaSessionStateFailedReason = 5 // Session cancelled.

)

type MMCallState uint32 // State of Call.

//go:generate stringer -type=MMCallState
const (
	MmCallStateUnknown    MMCallState = 0 // default state for a new outgoing call.
	MmCallStateDialing    MMCallState = 1 // outgoing call started. Wait for free channel.
	MmCallStateRingingOut MMCallState = 2 // incoming call is waiting for an answer.
	MmCallStateRingingIn  MMCallState = 3 // outgoing call attached to GSM network, waiting for an answer.
	MmCallStateActive     MMCallState = 4 // call is active between two peers.
	MmCallStateHeld       MMCallState = 5 // held call (by +CHLD AT command).
	MmCallStateWaiting    MMCallState = 6 // waiting call (by +CCWA AT command).
	MmCallStateTerminated MMCallState = 7 // call is terminated.

)

type MMCallStateReason uint32 // Reason for the state change in the call.

//go:generate stringer -type=MMCallStateReason
const (
	MmCallStateReasonUnknown          MMCallStateReason = 0 // Default value for a new outgoing call.
	MmCallStateReasonOutgoingStarted  MMCallStateReason = 1 // Outgoing call is started.
	MmCallStateReasonIncomingNew      MMCallStateReason = 2 // Received a new incoming call.
	MmCallStateReasonAccepted         MMCallStateReason = 3 // Dialing or Ringing call is accepted.
	MmCallStateReasonTerminated       MMCallStateReason = 4 // Call is correctly terminated.
	MmCallStateReasonRefusedOrBusy    MMCallStateReason = 5 // Remote peer is busy or refused call.
	MmCallStateReasonError            MMCallStateReason = 6 // Wrong number or generic network error.
	MmCallStateReasonAudioSetupFailed MMCallStateReason = 7 // Error setting up audio channel.
	MmCallStateReasonTransferred      MMCallStateReason = 8 // Call has been transferred.
	MmCallStateReasonDeflected        MMCallStateReason = 9 // Call has been deflected to a new number.
)

type MMCallDirection uint32 // Direction of the call.

//go:generate stringer -type=MMCallDirection
const (
	MmCallDirectionUnknown  MMCallDirection = 0 // unknown.
	MmCallDirectionIncoming MMCallDirection = 1 // call from network.
	MmCallDirectionOutgoing MMCallDirection = 2 // call to network.

)

type MMModemFirmwareUpdateMethod uint32 // Type of firmware update method supported by the module.

//go:generate stringer -type=MMModemFirmwareUpdateMethod
const (
	MmModemFirmwareUpdateMethodNone     MMModemFirmwareUpdateMethod = 0      // No method specified.
	MmModemFirmwareUpdateMethodFastboot MMModemFirmwareUpdateMethod = 1 << 0 // Device supports fastboot-based update.
	MmModemFirmwareUpdateMethodQmiPdc   MMModemFirmwareUpdateMethod = 1 << 1 // Device supports QMI PDC based update.

)