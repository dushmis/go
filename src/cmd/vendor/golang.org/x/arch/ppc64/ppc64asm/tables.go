package ppc64asm

const (
	_ Op = iota
	CNTLZW
	CNTLZW_
	B
	BA
	BL
	BLA
	BC
	BCA
	BCL
	BCLA
	BCLR
	BCLRL
	BCCTR
	BCCTRL
	BCTAR
	BCTARL
	CRAND
	CROR
	CRNAND
	CRXOR
	CRNOR
	CRANDC
	MCRF
	CREQV
	CRORC
	SC
	CLRBHRB
	MFBHRBE
	LBZ
	LBZU
	LBZX
	LBZUX
	LHZ
	LHZU
	LHZX
	LHZUX
	LHA
	LHAU
	LHAX
	LHAUX
	LWZ
	LWZU
	LWZX
	LWZUX
	LWA
	LWAX
	LWAUX
	LD
	LDU
	LDX
	LDUX
	STB
	STBU
	STBX
	STBUX
	STH
	STHU
	STHX
	STHUX
	STW
	STWU
	STWX
	STWUX
	STD
	STDU
	STDX
	STDUX
	LQ
	STQ
	LHBRX
	LWBRX
	STHBRX
	STWBRX
	LDBRX
	STDBRX
	LMW
	STMW
	LSWI
	LSWX
	STSWI
	STSWX
	LI
	ADDI
	LIS
	ADDIS
	ADD
	ADD_
	ADDO
	ADDO_
	ADDIC
	SUBF
	SUBF_
	SUBFO
	SUBFO_
	ADDIC_
	SUBFIC
	ADDC
	ADDC_
	ADDCO
	ADDCO_
	SUBFC
	SUBFC_
	SUBFCO
	SUBFCO_
	ADDE
	ADDE_
	ADDEO
	ADDEO_
	ADDME
	ADDME_
	ADDMEO
	ADDMEO_
	SUBFE
	SUBFE_
	SUBFEO
	SUBFEO_
	SUBFME
	SUBFME_
	SUBFMEO
	SUBFMEO_
	ADDZE
	ADDZE_
	ADDZEO
	ADDZEO_
	SUBFZE
	SUBFZE_
	SUBFZEO
	SUBFZEO_
	NEG
	NEG_
	NEGO
	NEGO_
	MULLI
	MULLW
	MULLW_
	MULLWO
	MULLWO_
	MULHW
	MULHW_
	MULHWU
	MULHWU_
	DIVW
	DIVW_
	DIVWO
	DIVWO_
	DIVWU
	DIVWU_
	DIVWUO
	DIVWUO_
	DIVWE
	DIVWE_
	DIVWEO
	DIVWEO_
	DIVWEU
	DIVWEU_
	DIVWEUO
	DIVWEUO_
	MULLD
	MULLD_
	MULLDO
	MULLDO_
	MULHDU
	MULHDU_
	MULHD
	MULHD_
	DIVD
	DIVD_
	DIVDO
	DIVDO_
	DIVDU
	DIVDU_
	DIVDUO
	DIVDUO_
	DIVDE
	DIVDE_
	DIVDEO
	DIVDEO_
	DIVDEU
	DIVDEU_
	DIVDEUO
	DIVDEUO_
	CMPWI
	CMPDI
	CMPW
	CMPD
	CMPLWI
	CMPLDI
	CMPLW
	CMPLD
	TWI
	TW
	TDI
	ISEL
	TD
	ANDI_
	ANDIS_
	ORI
	ORIS
	XORI
	XORIS
	AND
	AND_
	XOR
	XOR_
	NAND
	NAND_
	OR
	OR_
	NOR
	NOR_
	ANDC
	ANDC_
	EXTSB
	EXTSB_
	EQV
	EQV_
	ORC
	ORC_
	EXTSH
	EXTSH_
	CMPB
	POPCNTB
	POPCNTW
	PRTYD
	PRTYW
	EXTSW
	EXTSW_
	CNTLZD
	CNTLZD_
	POPCNTD
	BPERMD
	RLWINM
	RLWINM_
	RLWNM
	RLWNM_
	RLWIMI
	RLWIMI_
	RLDICL
	RLDICL_
	RLDICR
	RLDICR_
	RLDIC
	RLDIC_
	RLDCL
	RLDCL_
	RLDCR
	RLDCR_
	RLDIMI
	RLDIMI_
	SLW
	SLW_
	SRW
	SRW_
	SRAWI
	SRAWI_
	SRAW
	SRAW_
	SLD
	SLD_
	SRD
	SRD_
	SRADI
	SRADI_
	SRAD
	SRAD_
	CDTBCD
	CBCDTD
	ADDG6S
	MTSPR
	MFSPR
	MTCRF
	MFCR
	MTSLE
	MFVSRD
	MFVSRWZ
	MTVSRD
	MTVSRWA
	MTVSRWZ
	MTOCRF
	MFOCRF
	MCRXR
	MTDCRUX
	MFDCRUX
	LFS
	LFSU
	LFSX
	LFSUX
	LFD
	LFDU
	LFDX
	LFDUX
	LFIWAX
	LFIWZX
	STFS
	STFSU
	STFSX
	STFSUX
	STFD
	STFDU
	STFDX
	STFDUX
	STFIWX
	LFDP
	LFDPX
	STFDP
	STFDPX
	FMR
	FMR_
	FABS
	FABS_
	FNABS
	FNABS_
	FNEG
	FNEG_
	FCPSGN
	FCPSGN_
	FMRGEW
	FMRGOW
	FADD
	FADD_
	FADDS
	FADDS_
	FSUB
	FSUB_
	FSUBS
	FSUBS_
	FMUL
	FMUL_
	FMULS
	FMULS_
	FDIV
	FDIV_
	FDIVS
	FDIVS_
	FSQRT
	FSQRT_
	FSQRTS
	FSQRTS_
	FRE
	FRE_
	FRES
	FRES_
	FRSQRTE
	FRSQRTE_
	FRSQRTES
	FRSQRTES_
	FTDIV
	FTSQRT
	FMADD
	FMADD_
	FMADDS
	FMADDS_
	FMSUB
	FMSUB_
	FMSUBS
	FMSUBS_
	FNMADD
	FNMADD_
	FNMADDS
	FNMADDS_
	FNMSUB
	FNMSUB_
	FNMSUBS
	FNMSUBS_
	FRSP
	FRSP_
	FCTID
	FCTID_
	FCTIDZ
	FCTIDZ_
	FCTIDU
	FCTIDU_
	FCTIDUZ
	FCTIDUZ_
	FCTIW
	FCTIW_
	FCTIWZ
	FCTIWZ_
	FCTIWU
	FCTIWU_
	FCTIWUZ
	FCTIWUZ_
	FCFID
	FCFID_
	FCFIDU
	FCFIDU_
	FCFIDS
	FCFIDS_
	FCFIDUS
	FCFIDUS_
	FRIN
	FRIN_
	FRIZ
	FRIZ_
	FRIP
	FRIP_
	FRIM
	FRIM_
	FCMPU
	FCMPO
	FSEL
	FSEL_
	MFFS
	MFFS_
	MCRFS
	MTFSFI
	MTFSFI_
	MTFSF
	MTFSF_
	MTFSB0
	MTFSB0_
	MTFSB1
	MTFSB1_
	LVEBX
	LVEHX
	LVEWX
	LVX
	LVXL
	STVEBX
	STVEHX
	STVEWX
	STVX
	STVXL
	LVSL
	LVSR
	VPKPX
	VPKSDSS
	VPKSDUS
	VPKSHSS
	VPKSHUS
	VPKSWSS
	VPKSWUS
	VPKUDUM
	VPKUDUS
	VPKUHUM
	VPKUHUS
	VPKUWUM
	VPKUWUS
	VUPKHPX
	VUPKLPX
	VUPKHSB
	VUPKHSH
	VUPKHSW
	VUPKLSB
	VUPKLSH
	VUPKLSW
	VMRGHB
	VMRGHH
	VMRGLB
	VMRGLH
	VMRGHW
	VMRGLW
	VMRGEW
	VMRGOW
	VSPLTB
	VSPLTH
	VSPLTW
	VSPLTISB
	VSPLTISH
	VSPLTISW
	VPERM
	VSEL
	VSL
	VSLDOI
	VSLO
	VSR
	VSRO
	VADDCUW
	VADDSBS
	VADDSHS
	VADDSWS
	VADDUBM
	VADDUDM
	VADDUHM
	VADDUWM
	VADDUBS
	VADDUHS
	VADDUWS
	VADDUQM
	VADDEUQM
	VADDCUQ
	VADDECUQ
	VSUBCUW
	VSUBSBS
	VSUBSHS
	VSUBSWS
	VSUBUBM
	VSUBUDM
	VSUBUHM
	VSUBUWM
	VSUBUBS
	VSUBUHS
	VSUBUWS
	VSUBUQM
	VSUBEUQM
	VSUBCUQ
	VSUBECUQ
	VMULESB
	VMULEUB
	VMULOSB
	VMULOUB
	VMULESH
	VMULEUH
	VMULOSH
	VMULOUH
	VMULESW
	VMULEUW
	VMULOSW
	VMULOUW
	VMULUWM
	VMHADDSHS
	VMHRADDSHS
	VMLADDUHM
	VMSUMUBM
	VMSUMMBM
	VMSUMSHM
	VMSUMSHS
	VMSUMUHM
	VMSUMUHS
	VSUMSWS
	VSUM2SWS
	VSUM4SBS
	VSUM4SHS
	VSUM4UBS
	VAVGSB
	VAVGSH
	VAVGSW
	VAVGUB
	VAVGUW
	VAVGUH
	VMAXSB
	VMAXSD
	VMAXUB
	VMAXUD
	VMAXSH
	VMAXSW
	VMAXUH
	VMAXUW
	VMINSB
	VMINSD
	VMINUB
	VMINUD
	VMINSH
	VMINSW
	VMINUH
	VMINUW
	VCMPEQUB
	VCMPEQUB_
	VCMPEQUH
	VCMPEQUH_
	VCMPEQUW
	VCMPEQUW_
	VCMPEQUD
	VCMPEQUD_
	VCMPGTSB
	VCMPGTSB_
	VCMPGTSD
	VCMPGTSD_
	VCMPGTSH
	VCMPGTSH_
	VCMPGTSW
	VCMPGTSW_
	VCMPGTUB
	VCMPGTUB_
	VCMPGTUD
	VCMPGTUD_
	VCMPGTUH
	VCMPGTUH_
	VCMPGTUW
	VCMPGTUW_
	VAND
	VANDC
	VEQV
	VNAND
	VORC
	VNOR
	VOR
	VXOR
	VRLB
	VRLH
	VRLW
	VRLD
	VSLB
	VSLH
	VSLW
	VSLD
	VSRB
	VSRH
	VSRW
	VSRD
	VSRAB
	VSRAH
	VSRAW
	VSRAD
	VADDFP
	VSUBFP
	VMADDFP
	VNMSUBFP
	VMAXFP
	VMINFP
	VCTSXS
	VCTUXS
	VCFSX
	VCFUX
	VRFIM
	VRFIN
	VRFIP
	VRFIZ
	VCMPBFP
	VCMPBFP_
	VCMPEQFP
	VCMPEQFP_
	VCMPGEFP
	VCMPGEFP_
	VCMPGTFP
	VCMPGTFP_
	VEXPTEFP
	VLOGEFP
	VREFP
	VRSQRTEFP
	VCIPHER
	VCIPHERLAST
	VNCIPHER
	VNCIPHERLAST
	VSBOX
	VSHASIGMAD
	VSHASIGMAW
	VPMSUMB
	VPMSUMD
	VPMSUMH
	VPMSUMW
	VPERMXOR
	VGBBD
	VCLZB
	VCLZH
	VCLZW
	VCLZD
	VPOPCNTB
	VPOPCNTD
	VPOPCNTH
	VPOPCNTW
	VBPERMQ
	BCDADD_
	BCDSUB_
	MTVSCR
	MFVSCR
	DADD
	DADD_
	DSUB
	DSUB_
	DMUL
	DMUL_
	DDIV
	DDIV_
	DCMPU
	DCMPO
	DTSTDC
	DTSTDG
	DTSTEX
	DTSTSF
	DQUAI
	DQUAI_
	DQUA
	DQUA_
	DRRND
	DRRND_
	DRINTX
	DRINTX_
	DRINTN
	DRINTN_
	DCTDP
	DCTDP_
	DCTQPQ
	DCTQPQ_
	DRSP
	DRSP_
	DRDPQ
	DRDPQ_
	DCFFIX
	DCFFIX_
	DCFFIXQ
	DCFFIXQ_
	DCTFIX
	DCTFIX_
	DDEDPD
	DDEDPD_
	DENBCD
	DENBCD_
	DXEX
	DXEX_
	DIEX
	DIEX_
	DSCLI
	DSCLI_
	DSCRI
	DSCRI_
	LXSDX
	LXSIWAX
	LXSIWZX
	LXSSPX
	LXVD2X
	LXVDSX
	LXVW4X
	STXSDX
	STXSIWX
	STXSSPX
	STXVD2X
	STXVW4X
	XSABSDP
	XSADDDP
	XSADDSP
	XSCMPODP
	XSCMPUDP
	XSCPSGNDP
	XSCVDPSP
	XSCVDPSPN
	XSCVDPSXDS
	XSCVDPSXWS
	XSCVDPUXDS
	XSCVDPUXWS
	XSCVSPDP
	XSCVSPDPN
	XSCVSXDDP
	XSCVSXDSP
	XSCVUXDDP
	XSCVUXDSP
	XSDIVDP
	XSDIVSP
	XSMADDADP
	XSMADDASP
	XSMAXDP
	XSMINDP
	XSMSUBADP
	XSMSUBASP
	XSMULDP
	XSMULSP
	XSNABSDP
	XSNEGDP
	XSNMADDADP
	XSNMADDASP
	XSNMSUBADP
	XSNMSUBASP
	XSRDPI
	XSRDPIC
	XSRDPIM
	XSRDPIP
	XSRDPIZ
	XSREDP
	XSRESP
	XSRSP
	XSRSQRTEDP
	XSRSQRTESP
	XSSQRTDP
	XSSQRTSP
	XSSUBDP
	XSSUBSP
	XSTDIVDP
	XSTSQRTDP
	XVABSDP
	XVABSSP
	XVADDDP
	XVADDSP
	XVCMPEQDP
	XVCMPEQDP_
	XVCMPEQSP
	XVCMPEQSP_
	XVCMPGEDP
	XVCMPGEDP_
	XVCMPGESP
	XVCMPGESP_
	XVCMPGTDP
	XVCMPGTDP_
	XVCMPGTSP
	XVCMPGTSP_
	XVCPSGNDP
	XVCPSGNSP
	XVCVDPSP
	XVCVDPSXDS
	XVCVDPSXWS
	XVCVDPUXDS
	XVCVDPUXWS
	XVCVSPDP
	XVCVSPSXDS
	XVCVSPSXWS
	XVCVSPUXDS
	XVCVSPUXWS
	XVCVSXDDP
	XVCVSXDSP
	XVCVSXWDP
	XVCVSXWSP
	XVCVUXDDP
	XVCVUXDSP
	XVCVUXWDP
	XVCVUXWSP
	XVDIVDP
	XVDIVSP
	XVMADDADP
	XVMADDASP
	XVMAXDP
	XVMAXSP
	XVMINDP
	XVMINSP
	XVMSUBADP
	XVMSUBASP
	XVMULDP
	XVMULSP
	XVNABSDP
	XVNABSSP
	XVNEGDP
	XVNEGSP
	XVNMADDADP
	XVNMADDASP
	XVNMSUBADP
	XVNMSUBASP
	XVRDPI
	XVRDPIC
	XVRDPIM
	XVRDPIP
	XVRDPIZ
	XVREDP
	XVRESP
	XVRSPI
	XVRSPIC
	XVRSPIM
	XVRSPIP
	XVRSPIZ
	XVRSQRTEDP
	XVRSQRTESP
	XVSQRTDP
	XVSQRTSP
	XVSUBDP
	XVSUBSP
	XVTDIVDP
	XVTDIVSP
	XVTSQRTDP
	XVTSQRTSP
	XXLAND
	XXLANDC
	XXLEQV
	XXLNAND
	XXLORC
	XXLNOR
	XXLOR
	XXLXOR
	XXMRGHW
	XXMRGLW
	XXPERMDI
	XXSEL
	XXSLDWI
	XXSPLTW
	BRINC
	EVABS
	EVADDIW
	EVADDSMIAAW
	EVADDSSIAAW
	EVADDUMIAAW
	EVADDUSIAAW
	EVADDW
	EVAND
	EVCMPEQ
	EVANDC
	EVCMPGTS
	EVCMPGTU
	EVCMPLTU
	EVCMPLTS
	EVCNTLSW
	EVCNTLZW
	EVDIVWS
	EVDIVWU
	EVEQV
	EVEXTSB
	EVEXTSH
	EVLDD
	EVLDH
	EVLDDX
	EVLDHX
	EVLDW
	EVLHHESPLAT
	EVLDWX
	EVLHHESPLATX
	EVLHHOSSPLAT
	EVLHHOUSPLAT
	EVLHHOSSPLATX
	EVLHHOUSPLATX
	EVLWHE
	EVLWHOS
	EVLWHEX
	EVLWHOSX
	EVLWHOU
	EVLWHSPLAT
	EVLWHOUX
	EVLWHSPLATX
	EVLWWSPLAT
	EVMERGEHI
	EVLWWSPLATX
	EVMERGELO
	EVMERGEHILO
	EVMHEGSMFAA
	EVMERGELOHI
	EVMHEGSMFAN
	EVMHEGSMIAA
	EVMHEGUMIAA
	EVMHEGSMIAN
	EVMHEGUMIAN
	EVMHESMF
	EVMHESMFAAW
	EVMHESMFA
	EVMHESMFANW
	EVMHESMI
	EVMHESMIAAW
	EVMHESMIA
	EVMHESMIANW
	EVMHESSF
	EVMHESSFA
	EVMHESSFAAW
	EVMHESSFANW
	EVMHESSIAAW
	EVMHESSIANW
	EVMHEUMI
	EVMHEUMIAAW
	EVMHEUMIA
	EVMHEUMIANW
	EVMHEUSIAAW
	EVMHEUSIANW
	EVMHOGSMFAA
	EVMHOGSMIAA
	EVMHOGSMFAN
	EVMHOGSMIAN
	EVMHOGUMIAA
	EVMHOSMF
	EVMHOGUMIAN
	EVMHOSMFA
	EVMHOSMFAAW
	EVMHOSMI
	EVMHOSMFANW
	EVMHOSMIA
	EVMHOSMIAAW
	EVMHOSMIANW
	EVMHOSSF
	EVMHOSSFA
	EVMHOSSFAAW
	EVMHOSSFANW
	EVMHOSSIAAW
	EVMHOUMI
	EVMHOSSIANW
	EVMHOUMIA
	EVMHOUMIAAW
	EVMHOUSIAAW
	EVMHOUMIANW
	EVMHOUSIANW
	EVMRA
	EVMWHSMF
	EVMWHSMI
	EVMWHSMFA
	EVMWHSMIA
	EVMWHSSF
	EVMWHUMI
	EVMWHSSFA
	EVMWHUMIA
	EVMWLSMIAAW
	EVMWLSSIAAW
	EVMWLSMIANW
	EVMWLSSIANW
	EVMWLUMI
	EVMWLUMIAAW
	EVMWLUMIA
	EVMWLUMIANW
	EVMWLUSIAAW
	EVMWSMF
	EVMWLUSIANW
	EVMWSMFA
	EVMWSMFAA
	EVMWSMI
	EVMWSMIAA
	EVMWSMFAN
	EVMWSMIA
	EVMWSMIAN
	EVMWSSF
	EVMWSSFA
	EVMWSSFAA
	EVMWUMI
	EVMWSSFAN
	EVMWUMIA
	EVMWUMIAA
	EVNAND
	EVMWUMIAN
	EVNEG
	EVNOR
	EVORC
	EVOR
	EVRLW
	EVRLWI
	EVSEL
	EVRNDW
	EVSLW
	EVSPLATFI
	EVSRWIS
	EVSLWI
	EVSPLATI
	EVSRWIU
	EVSRWS
	EVSTDD
	EVSRWU
	EVSTDDX
	EVSTDH
	EVSTDW
	EVSTDHX
	EVSTDWX
	EVSTWHE
	EVSTWHO
	EVSTWWE
	EVSTWHEX
	EVSTWHOX
	EVSTWWEX
	EVSTWWO
	EVSUBFSMIAAW
	EVSTWWOX
	EVSUBFSSIAAW
	EVSUBFUMIAAW
	EVSUBFUSIAAW
	EVSUBFW
	EVSUBIFW
	EVXOR
	EVFSABS
	EVFSNABS
	EVFSNEG
	EVFSADD
	EVFSMUL
	EVFSSUB
	EVFSDIV
	EVFSCMPGT
	EVFSCMPLT
	EVFSCMPEQ
	EVFSTSTGT
	EVFSTSTLT
	EVFSTSTEQ
	EVFSCFSI
	EVFSCFSF
	EVFSCFUI
	EVFSCFUF
	EVFSCTSI
	EVFSCTUI
	EVFSCTSIZ
	EVFSCTUIZ
	EVFSCTSF
	EVFSCTUF
	EFSABS
	EFSNEG
	EFSNABS
	EFSADD
	EFSMUL
	EFSSUB
	EFSDIV
	EFSCMPGT
	EFSCMPLT
	EFSCMPEQ
	EFSTSTGT
	EFSTSTLT
	EFSTSTEQ
	EFSCFSI
	EFSCFSF
	EFSCTSI
	EFSCFUI
	EFSCFUF
	EFSCTUI
	EFSCTSIZ
	EFSCTSF
	EFSCTUIZ
	EFSCTUF
	EFDABS
	EFDNEG
	EFDNABS
	EFDADD
	EFDMUL
	EFDSUB
	EFDDIV
	EFDCMPGT
	EFDCMPEQ
	EFDCMPLT
	EFDTSTGT
	EFDTSTLT
	EFDCFSI
	EFDTSTEQ
	EFDCFUI
	EFDCFSID
	EFDCFSF
	EFDCFUF
	EFDCFUID
	EFDCTSI
	EFDCTUI
	EFDCTSIDZ
	EFDCTUIDZ
	EFDCTSIZ
	EFDCTSF
	EFDCTUF
	EFDCTUIZ
	EFDCFS
	EFSCFD
	DLMZB
	DLMZB_
	MACCHW
	MACCHW_
	MACCHWO
	MACCHWO_
	MACCHWS
	MACCHWS_
	MACCHWSO
	MACCHWSO_
	MACCHWU
	MACCHWU_
	MACCHWUO
	MACCHWUO_
	MACCHWSU
	MACCHWSU_
	MACCHWSUO
	MACCHWSUO_
	MACHHW
	MACHHW_
	MACHHWO
	MACHHWO_
	MACHHWS
	MACHHWS_
	MACHHWSO
	MACHHWSO_
	MACHHWU
	MACHHWU_
	MACHHWUO
	MACHHWUO_
	MACHHWSU
	MACHHWSU_
	MACHHWSUO
	MACHHWSUO_
	MACLHW
	MACLHW_
	MACLHWO
	MACLHWO_
	MACLHWS
	MACLHWS_
	MACLHWSO
	MACLHWSO_
	MACLHWU
	MACLHWU_
	MACLHWUO
	MACLHWUO_
	MULCHW
	MULCHW_
	MACLHWSU
	MACLHWSU_
	MACLHWSUO
	MACLHWSUO_
	MULCHWU
	MULCHWU_
	MULHHW
	MULHHW_
	MULLHW
	MULLHW_
	MULHHWU
	MULHHWU_
	MULLHWU
	MULLHWU_
	NMACCHW
	NMACCHW_
	NMACCHWO
	NMACCHWO_
	NMACCHWS
	NMACCHWS_
	NMACCHWSO
	NMACCHWSO_
	NMACHHW
	NMACHHW_
	NMACHHWO
	NMACHHWO_
	NMACHHWS
	NMACHHWS_
	NMACHHWSO
	NMACHHWSO_
	NMACLHW
	NMACLHW_
	NMACLHWO
	NMACLHWO_
	NMACLHWS
	NMACLHWS_
	NMACLHWSO
	NMACLHWSO_
	ICBI
	ICBT
	DCBA
	DCBT
	DCBTST
	DCBZ
	DCBST
	DCBF
	ISYNC
	LBARX
	LHARX
	LWARX
	STBCX_
	STHCX_
	STWCX_
	LDARX
	STDCX_
	LQARX
	STQCX_
	SYNC
	EIEIO
	MBAR
	WAIT
	TBEGIN_
	TEND_
	TABORT_
	TABORTWC_
	TABORTWCI_
	TABORTDC_
	TABORTDCI_
	TSR_
	TCHECK
	MFTB
	RFEBB
	LBDX
	LHDX
	LWDX
	LDDX
	LFDDX
	STBDX
	STHDX
	STWDX
	STDDX
	STFDDX
	DSN
	ECIWX
	ECOWX
	RFID
	HRFID
	DOZE
	NAP
	SLEEP
	RVWINKLE
	LBZCIX
	LWZCIX
	LHZCIX
	LDCIX
	STBCIX
	STWCIX
	STHCIX
	STDCIX
	TRECLAIM_
	TRECHKPT_
	MTMSR
	MTMSRD
	MFMSR
	SLBIE
	SLBIA
	SLBMTE
	SLBMFEV
	SLBMFEE
	SLBFEE_
	MTSR
	MTSRIN
	MFSR
	MFSRIN
	TLBIE
	TLBIEL
	TLBIA
	TLBSYNC
	MSGSND
	MSGCLR
	MSGSNDP
	MSGCLRP
	MTTMR
	RFI
	RFCI
	RFDI
	RFMCI
	RFGI
	EHPRIV
	MTDCR
	MTDCRX
	MFDCR
	MFDCRX
	WRTEE
	WRTEEI
	LBEPX
	LHEPX
	LWEPX
	LDEPX
	STBEPX
	STHEPX
	STWEPX
	STDEPX
	DCBSTEP
	DCBTEP
	DCBFEP
	DCBTSTEP
	ICBIEP
	DCBZEP
	LFDEPX
	STFDEPX
	EVLDDEPX
	EVSTDDEPX
	LVEPX
	LVEPXL
	STVEPX
	STVEPXL
	DCBI
	DCBLQ_
	ICBLQ_
	DCBTLS
	DCBTSTLS
	ICBTLS
	ICBLC
	DCBLC
	TLBIVAX
	TLBILX
	TLBSX
	TLBSRX_
	TLBRE
	TLBWE
	DNH
	DCI
	ICI
	DCREAD
	ICREAD
	MFPMR
	MTPMR
)

var opstr = [...]string{
	CNTLZW:        "cntlzw",
	CNTLZW_:       "cntlzw.",
	B:             "b",
	BA:            "ba",
	BL:            "bl",
	BLA:           "bla",
	BC:            "bc",
	BCA:           "bca",
	BCL:           "bcl",
	BCLA:          "bcla",
	BCLR:          "bclr",
	BCLRL:         "bclrl",
	BCCTR:         "bcctr",
	BCCTRL:        "bcctrl",
	BCTAR:         "bctar",
	BCTARL:        "bctarl",
	CRAND:         "crand",
	CROR:          "cror",
	CRNAND:        "crnand",
	CRXOR:         "crxor",
	CRNOR:         "crnor",
	CRANDC:        "crandc",
	MCRF:          "mcrf",
	CREQV:         "creqv",
	CRORC:         "crorc",
	SC:            "sc",
	CLRBHRB:       "clrbhrb",
	MFBHRBE:       "mfbhrbe",
	LBZ:           "lbz",
	LBZU:          "lbzu",
	LBZX:          "lbzx",
	LBZUX:         "lbzux",
	LHZ:           "lhz",
	LHZU:          "lhzu",
	LHZX:          "lhzx",
	LHZUX:         "lhzux",
	LHA:           "lha",
	LHAU:          "lhau",
	LHAX:          "lhax",
	LHAUX:         "lhaux",
	LWZ:           "lwz",
	LWZU:          "lwzu",
	LWZX:          "lwzx",
	LWZUX:         "lwzux",
	LWA:           "lwa",
	LWAX:          "lwax",
	LWAUX:         "lwaux",
	LD:            "ld",
	LDU:           "ldu",
	LDX:           "ldx",
	LDUX:          "ldux",
	STB:           "stb",
	STBU:          "stbu",
	STBX:          "stbx",
	STBUX:         "stbux",
	STH:           "sth",
	STHU:          "sthu",
	STHX:          "sthx",
	STHUX:         "sthux",
	STW:           "stw",
	STWU:          "stwu",
	STWX:          "stwx",
	STWUX:         "stwux",
	STD:           "std",
	STDU:          "stdu",
	STDX:          "stdx",
	STDUX:         "stdux",
	LQ:            "lq",
	STQ:           "stq",
	LHBRX:         "lhbrx",
	LWBRX:         "lwbrx",
	STHBRX:        "sthbrx",
	STWBRX:        "stwbrx",
	LDBRX:         "ldbrx",
	STDBRX:        "stdbrx",
	LMW:           "lmw",
	STMW:          "stmw",
	LSWI:          "lswi",
	LSWX:          "lswx",
	STSWI:         "stswi",
	STSWX:         "stswx",
	LI:            "li",
	ADDI:          "addi",
	LIS:           "lis",
	ADDIS:         "addis",
	ADD:           "add",
	ADD_:          "add.",
	ADDO:          "addo",
	ADDO_:         "addo.",
	ADDIC:         "addic",
	SUBF:          "subf",
	SUBF_:         "subf.",
	SUBFO:         "subfo",
	SUBFO_:        "subfo.",
	ADDIC_:        "addic.",
	SUBFIC:        "subfic",
	ADDC:          "addc",
	ADDC_:         "addc.",
	ADDCO:         "addco",
	ADDCO_:        "addco.",
	SUBFC:         "subfc",
	SUBFC_:        "subfc.",
	SUBFCO:        "subfco",
	SUBFCO_:       "subfco.",
	ADDE:          "adde",
	ADDE_:         "adde.",
	ADDEO:         "addeo",
	ADDEO_:        "addeo.",
	ADDME:         "addme",
	ADDME_:        "addme.",
	ADDMEO:        "addmeo",
	ADDMEO_:       "addmeo.",
	SUBFE:         "subfe",
	SUBFE_:        "subfe.",
	SUBFEO:        "subfeo",
	SUBFEO_:       "subfeo.",
	SUBFME:        "subfme",
	SUBFME_:       "subfme.",
	SUBFMEO:       "subfmeo",
	SUBFMEO_:      "subfmeo.",
	ADDZE:         "addze",
	ADDZE_:        "addze.",
	ADDZEO:        "addzeo",
	ADDZEO_:       "addzeo.",
	SUBFZE:        "subfze",
	SUBFZE_:       "subfze.",
	SUBFZEO:       "subfzeo",
	SUBFZEO_:      "subfzeo.",
	NEG:           "neg",
	NEG_:          "neg.",
	NEGO:          "nego",
	NEGO_:         "nego.",
	MULLI:         "mulli",
	MULLW:         "mullw",
	MULLW_:        "mullw.",
	MULLWO:        "mullwo",
	MULLWO_:       "mullwo.",
	MULHW:         "mulhw",
	MULHW_:        "mulhw.",
	MULHWU:        "mulhwu",
	MULHWU_:       "mulhwu.",
	DIVW:          "divw",
	DIVW_:         "divw.",
	DIVWO:         "divwo",
	DIVWO_:        "divwo.",
	DIVWU:         "divwu",
	DIVWU_:        "divwu.",
	DIVWUO:        "divwuo",
	DIVWUO_:       "divwuo.",
	DIVWE:         "divwe",
	DIVWE_:        "divwe.",
	DIVWEO:        "divweo",
	DIVWEO_:       "divweo.",
	DIVWEU:        "divweu",
	DIVWEU_:       "divweu.",
	DIVWEUO:       "divweuo",
	DIVWEUO_:      "divweuo.",
	MULLD:         "mulld",
	MULLD_:        "mulld.",
	MULLDO:        "mulldo",
	MULLDO_:       "mulldo.",
	MULHDU:        "mulhdu",
	MULHDU_:       "mulhdu.",
	MULHD:         "mulhd",
	MULHD_:        "mulhd.",
	DIVD:          "divd",
	DIVD_:         "divd.",
	DIVDO:         "divdo",
	DIVDO_:        "divdo.",
	DIVDU:         "divdu",
	DIVDU_:        "divdu.",
	DIVDUO:        "divduo",
	DIVDUO_:       "divduo.",
	DIVDE:         "divde",
	DIVDE_:        "divde.",
	DIVDEO:        "divdeo",
	DIVDEO_:       "divdeo.",
	DIVDEU:        "divdeu",
	DIVDEU_:       "divdeu.",
	DIVDEUO:       "divdeuo",
	DIVDEUO_:      "divdeuo.",
	CMPWI:         "cmpwi",
	CMPDI:         "cmpdi",
	CMPW:          "cmpw",
	CMPD:          "cmpd",
	CMPLWI:        "cmplwi",
	CMPLDI:        "cmpldi",
	CMPLW:         "cmplw",
	CMPLD:         "cmpld",
	TWI:           "twi",
	TW:            "tw",
	TDI:           "tdi",
	ISEL:          "isel",
	TD:            "td",
	ANDI_:         "andi.",
	ANDIS_:        "andis.",
	ORI:           "ori",
	ORIS:          "oris",
	XORI:          "xori",
	XORIS:         "xoris",
	AND:           "and",
	AND_:          "and.",
	XOR:           "xor",
	XOR_:          "xor.",
	NAND:          "nand",
	NAND_:         "nand.",
	OR:            "or",
	OR_:           "or.",
	NOR:           "nor",
	NOR_:          "nor.",
	ANDC:          "andc",
	ANDC_:         "andc.",
	EXTSB:         "extsb",
	EXTSB_:        "extsb.",
	EQV:           "eqv",
	EQV_:          "eqv.",
	ORC:           "orc",
	ORC_:          "orc.",
	EXTSH:         "extsh",
	EXTSH_:        "extsh.",
	CMPB:          "cmpb",
	POPCNTB:       "popcntb",
	POPCNTW:       "popcntw",
	PRTYD:         "prtyd",
	PRTYW:         "prtyw",
	EXTSW:         "extsw",
	EXTSW_:        "extsw.",
	CNTLZD:        "cntlzd",
	CNTLZD_:       "cntlzd.",
	POPCNTD:       "popcntd",
	BPERMD:        "bpermd",
	RLWINM:        "rlwinm",
	RLWINM_:       "rlwinm.",
	RLWNM:         "rlwnm",
	RLWNM_:        "rlwnm.",
	RLWIMI:        "rlwimi",
	RLWIMI_:       "rlwimi.",
	RLDICL:        "rldicl",
	RLDICL_:       "rldicl.",
	RLDICR:        "rldicr",
	RLDICR_:       "rldicr.",
	RLDIC:         "rldic",
	RLDIC_:        "rldic.",
	RLDCL:         "rldcl",
	RLDCL_:        "rldcl.",
	RLDCR:         "rldcr",
	RLDCR_:        "rldcr.",
	RLDIMI:        "rldimi",
	RLDIMI_:       "rldimi.",
	SLW:           "slw",
	SLW_:          "slw.",
	SRW:           "srw",
	SRW_:          "srw.",
	SRAWI:         "srawi",
	SRAWI_:        "srawi.",
	SRAW:          "sraw",
	SRAW_:         "sraw.",
	SLD:           "sld",
	SLD_:          "sld.",
	SRD:           "srd",
	SRD_:          "srd.",
	SRADI:         "sradi",
	SRADI_:        "sradi.",
	SRAD:          "srad",
	SRAD_:         "srad.",
	CDTBCD:        "cdtbcd",
	CBCDTD:        "cbcdtd",
	ADDG6S:        "addg6s",
	MTSPR:         "mtspr",
	MFSPR:         "mfspr",
	MTCRF:         "mtcrf",
	MFCR:          "mfcr",
	MTSLE:         "mtsle",
	MFVSRD:        "mfvsrd",
	MFVSRWZ:       "mfvsrwz",
	MTVSRD:        "mtvsrd",
	MTVSRWA:       "mtvsrwa",
	MTVSRWZ:       "mtvsrwz",
	MTOCRF:        "mtocrf",
	MFOCRF:        "mfocrf",
	MCRXR:         "mcrxr",
	MTDCRUX:       "mtdcrux",
	MFDCRUX:       "mfdcrux",
	LFS:           "lfs",
	LFSU:          "lfsu",
	LFSX:          "lfsx",
	LFSUX:         "lfsux",
	LFD:           "lfd",
	LFDU:          "lfdu",
	LFDX:          "lfdx",
	LFDUX:         "lfdux",
	LFIWAX:        "lfiwax",
	LFIWZX:        "lfiwzx",
	STFS:          "stfs",
	STFSU:         "stfsu",
	STFSX:         "stfsx",
	STFSUX:        "stfsux",
	STFD:          "stfd",
	STFDU:         "stfdu",
	STFDX:         "stfdx",
	STFDUX:        "stfdux",
	STFIWX:        "stfiwx",
	LFDP:          "lfdp",
	LFDPX:         "lfdpx",
	STFDP:         "stfdp",
	STFDPX:        "stfdpx",
	FMR:           "fmr",
	FMR_:          "fmr.",
	FABS:          "fabs",
	FABS_:         "fabs.",
	FNABS:         "fnabs",
	FNABS_:        "fnabs.",
	FNEG:          "fneg",
	FNEG_:         "fneg.",
	FCPSGN:        "fcpsgn",
	FCPSGN_:       "fcpsgn.",
	FMRGEW:        "fmrgew",
	FMRGOW:        "fmrgow",
	FADD:          "fadd",
	FADD_:         "fadd.",
	FADDS:         "fadds",
	FADDS_:        "fadds.",
	FSUB:          "fsub",
	FSUB_:         "fsub.",
	FSUBS:         "fsubs",
	FSUBS_:        "fsubs.",
	FMUL:          "fmul",
	FMUL_:         "fmul.",
	FMULS:         "fmuls",
	FMULS_:        "fmuls.",
	FDIV:          "fdiv",
	FDIV_:         "fdiv.",
	FDIVS:         "fdivs",
	FDIVS_:        "fdivs.",
	FSQRT:         "fsqrt",
	FSQRT_:        "fsqrt.",
	FSQRTS:        "fsqrts",
	FSQRTS_:       "fsqrts.",
	FRE:           "fre",
	FRE_:          "fre.",
	FRES:          "fres",
	FRES_:         "fres.",
	FRSQRTE:       "frsqrte",
	FRSQRTE_:      "frsqrte.",
	FRSQRTES:      "frsqrtes",
	FRSQRTES_:     "frsqrtes.",
	FTDIV:         "ftdiv",
	FTSQRT:        "ftsqrt",
	FMADD:         "fmadd",
	FMADD_:        "fmadd.",
	FMADDS:        "fmadds",
	FMADDS_:       "fmadds.",
	FMSUB:         "fmsub",
	FMSUB_:        "fmsub.",
	FMSUBS:        "fmsubs",
	FMSUBS_:       "fmsubs.",
	FNMADD:        "fnmadd",
	FNMADD_:       "fnmadd.",
	FNMADDS:       "fnmadds",
	FNMADDS_:      "fnmadds.",
	FNMSUB:        "fnmsub",
	FNMSUB_:       "fnmsub.",
	FNMSUBS:       "fnmsubs",
	FNMSUBS_:      "fnmsubs.",
	FRSP:          "frsp",
	FRSP_:         "frsp.",
	FCTID:         "fctid",
	FCTID_:        "fctid.",
	FCTIDZ:        "fctidz",
	FCTIDZ_:       "fctidz.",
	FCTIDU:        "fctidu",
	FCTIDU_:       "fctidu.",
	FCTIDUZ:       "fctiduz",
	FCTIDUZ_:      "fctiduz.",
	FCTIW:         "fctiw",
	FCTIW_:        "fctiw.",
	FCTIWZ:        "fctiwz",
	FCTIWZ_:       "fctiwz.",
	FCTIWU:        "fctiwu",
	FCTIWU_:       "fctiwu.",
	FCTIWUZ:       "fctiwuz",
	FCTIWUZ_:      "fctiwuz.",
	FCFID:         "fcfid",
	FCFID_:        "fcfid.",
	FCFIDU:        "fcfidu",
	FCFIDU_:       "fcfidu.",
	FCFIDS:        "fcfids",
	FCFIDS_:       "fcfids.",
	FCFIDUS:       "fcfidus",
	FCFIDUS_:      "fcfidus.",
	FRIN:          "frin",
	FRIN_:         "frin.",
	FRIZ:          "friz",
	FRIZ_:         "friz.",
	FRIP:          "frip",
	FRIP_:         "frip.",
	FRIM:          "frim",
	FRIM_:         "frim.",
	FCMPU:         "fcmpu",
	FCMPO:         "fcmpo",
	FSEL:          "fsel",
	FSEL_:         "fsel.",
	MFFS:          "mffs",
	MFFS_:         "mffs.",
	MCRFS:         "mcrfs",
	MTFSFI:        "mtfsfi",
	MTFSFI_:       "mtfsfi.",
	MTFSF:         "mtfsf",
	MTFSF_:        "mtfsf.",
	MTFSB0:        "mtfsb0",
	MTFSB0_:       "mtfsb0.",
	MTFSB1:        "mtfsb1",
	MTFSB1_:       "mtfsb1.",
	LVEBX:         "lvebx",
	LVEHX:         "lvehx",
	LVEWX:         "lvewx",
	LVX:           "lvx",
	LVXL:          "lvxl",
	STVEBX:        "stvebx",
	STVEHX:        "stvehx",
	STVEWX:        "stvewx",
	STVX:          "stvx",
	STVXL:         "stvxl",
	LVSL:          "lvsl",
	LVSR:          "lvsr",
	VPKPX:         "vpkpx",
	VPKSDSS:       "vpksdss",
	VPKSDUS:       "vpksdus",
	VPKSHSS:       "vpkshss",
	VPKSHUS:       "vpkshus",
	VPKSWSS:       "vpkswss",
	VPKSWUS:       "vpkswus",
	VPKUDUM:       "vpkudum",
	VPKUDUS:       "vpkudus",
	VPKUHUM:       "vpkuhum",
	VPKUHUS:       "vpkuhus",
	VPKUWUM:       "vpkuwum",
	VPKUWUS:       "vpkuwus",
	VUPKHPX:       "vupkhpx",
	VUPKLPX:       "vupklpx",
	VUPKHSB:       "vupkhsb",
	VUPKHSH:       "vupkhsh",
	VUPKHSW:       "vupkhsw",
	VUPKLSB:       "vupklsb",
	VUPKLSH:       "vupklsh",
	VUPKLSW:       "vupklsw",
	VMRGHB:        "vmrghb",
	VMRGHH:        "vmrghh",
	VMRGLB:        "vmrglb",
	VMRGLH:        "vmrglh",
	VMRGHW:        "vmrghw",
	VMRGLW:        "vmrglw",
	VMRGEW:        "vmrgew",
	VMRGOW:        "vmrgow",
	VSPLTB:        "vspltb",
	VSPLTH:        "vsplth",
	VSPLTW:        "vspltw",
	VSPLTISB:      "vspltisb",
	VSPLTISH:      "vspltish",
	VSPLTISW:      "vspltisw",
	VPERM:         "vperm",
	VSEL:          "vsel",
	VSL:           "vsl",
	VSLDOI:        "vsldoi",
	VSLO:          "vslo",
	VSR:           "vsr",
	VSRO:          "vsro",
	VADDCUW:       "vaddcuw",
	VADDSBS:       "vaddsbs",
	VADDSHS:       "vaddshs",
	VADDSWS:       "vaddsws",
	VADDUBM:       "vaddubm",
	VADDUDM:       "vaddudm",
	VADDUHM:       "vadduhm",
	VADDUWM:       "vadduwm",
	VADDUBS:       "vaddubs",
	VADDUHS:       "vadduhs",
	VADDUWS:       "vadduws",
	VADDUQM:       "vadduqm",
	VADDEUQM:      "vaddeuqm",
	VADDCUQ:       "vaddcuq",
	VADDECUQ:      "vaddecuq",
	VSUBCUW:       "vsubcuw",
	VSUBSBS:       "vsubsbs",
	VSUBSHS:       "vsubshs",
	VSUBSWS:       "vsubsws",
	VSUBUBM:       "vsububm",
	VSUBUDM:       "vsubudm",
	VSUBUHM:       "vsubuhm",
	VSUBUWM:       "vsubuwm",
	VSUBUBS:       "vsububs",
	VSUBUHS:       "vsubuhs",
	VSUBUWS:       "vsubuws",
	VSUBUQM:       "vsubuqm",
	VSUBEUQM:      "vsubeuqm",
	VSUBCUQ:       "vsubcuq",
	VSUBECUQ:      "vsubecuq",
	VMULESB:       "vmulesb",
	VMULEUB:       "vmuleub",
	VMULOSB:       "vmulosb",
	VMULOUB:       "vmuloub",
	VMULESH:       "vmulesh",
	VMULEUH:       "vmuleuh",
	VMULOSH:       "vmulosh",
	VMULOUH:       "vmulouh",
	VMULESW:       "vmulesw",
	VMULEUW:       "vmuleuw",
	VMULOSW:       "vmulosw",
	VMULOUW:       "vmulouw",
	VMULUWM:       "vmuluwm",
	VMHADDSHS:     "vmhaddshs",
	VMHRADDSHS:    "vmhraddshs",
	VMLADDUHM:     "vmladduhm",
	VMSUMUBM:      "vmsumubm",
	VMSUMMBM:      "vmsummbm",
	VMSUMSHM:      "vmsumshm",
	VMSUMSHS:      "vmsumshs",
	VMSUMUHM:      "vmsumuhm",
	VMSUMUHS:      "vmsumuhs",
	VSUMSWS:       "vsumsws",
	VSUM2SWS:      "vsum2sws",
	VSUM4SBS:      "vsum4sbs",
	VSUM4SHS:      "vsum4shs",
	VSUM4UBS:      "vsum4ubs",
	VAVGSB:        "vavgsb",
	VAVGSH:        "vavgsh",
	VAVGSW:        "vavgsw",
	VAVGUB:        "vavgub",
	VAVGUW:        "vavguw",
	VAVGUH:        "vavguh",
	VMAXSB:        "vmaxsb",
	VMAXSD:        "vmaxsd",
	VMAXUB:        "vmaxub",
	VMAXUD:        "vmaxud",
	VMAXSH:        "vmaxsh",
	VMAXSW:        "vmaxsw",
	VMAXUH:        "vmaxuh",
	VMAXUW:        "vmaxuw",
	VMINSB:        "vminsb",
	VMINSD:        "vminsd",
	VMINUB:        "vminub",
	VMINUD:        "vminud",
	VMINSH:        "vminsh",
	VMINSW:        "vminsw",
	VMINUH:        "vminuh",
	VMINUW:        "vminuw",
	VCMPEQUB:      "vcmpequb",
	VCMPEQUB_:     "vcmpequb.",
	VCMPEQUH:      "vcmpequh",
	VCMPEQUH_:     "vcmpequh.",
	VCMPEQUW:      "vcmpequw",
	VCMPEQUW_:     "vcmpequw.",
	VCMPEQUD:      "vcmpequd",
	VCMPEQUD_:     "vcmpequd.",
	VCMPGTSB:      "vcmpgtsb",
	VCMPGTSB_:     "vcmpgtsb.",
	VCMPGTSD:      "vcmpgtsd",
	VCMPGTSD_:     "vcmpgtsd.",
	VCMPGTSH:      "vcmpgtsh",
	VCMPGTSH_:     "vcmpgtsh.",
	VCMPGTSW:      "vcmpgtsw",
	VCMPGTSW_:     "vcmpgtsw.",
	VCMPGTUB:      "vcmpgtub",
	VCMPGTUB_:     "vcmpgtub.",
	VCMPGTUD:      "vcmpgtud",
	VCMPGTUD_:     "vcmpgtud.",
	VCMPGTUH:      "vcmpgtuh",
	VCMPGTUH_:     "vcmpgtuh.",
	VCMPGTUW:      "vcmpgtuw",
	VCMPGTUW_:     "vcmpgtuw.",
	VAND:          "vand",
	VANDC:         "vandc",
	VEQV:          "veqv",
	VNAND:         "vnand",
	VORC:          "vorc",
	VNOR:          "vnor",
	VOR:           "vor",
	VXOR:          "vxor",
	VRLB:          "vrlb",
	VRLH:          "vrlh",
	VRLW:          "vrlw",
	VRLD:          "vrld",
	VSLB:          "vslb",
	VSLH:          "vslh",
	VSLW:          "vslw",
	VSLD:          "vsld",
	VSRB:          "vsrb",
	VSRH:          "vsrh",
	VSRW:          "vsrw",
	VSRD:          "vsrd",
	VSRAB:         "vsrab",
	VSRAH:         "vsrah",
	VSRAW:         "vsraw",
	VSRAD:         "vsrad",
	VADDFP:        "vaddfp",
	VSUBFP:        "vsubfp",
	VMADDFP:       "vmaddfp",
	VNMSUBFP:      "vnmsubfp",
	VMAXFP:        "vmaxfp",
	VMINFP:        "vminfp",
	VCTSXS:        "vctsxs",
	VCTUXS:        "vctuxs",
	VCFSX:         "vcfsx",
	VCFUX:         "vcfux",
	VRFIM:         "vrfim",
	VRFIN:         "vrfin",
	VRFIP:         "vrfip",
	VRFIZ:         "vrfiz",
	VCMPBFP:       "vcmpbfp",
	VCMPBFP_:      "vcmpbfp.",
	VCMPEQFP:      "vcmpeqfp",
	VCMPEQFP_:     "vcmpeqfp.",
	VCMPGEFP:      "vcmpgefp",
	VCMPGEFP_:     "vcmpgefp.",
	VCMPGTFP:      "vcmpgtfp",
	VCMPGTFP_:     "vcmpgtfp.",
	VEXPTEFP:      "vexptefp",
	VLOGEFP:       "vlogefp",
	VREFP:         "vrefp",
	VRSQRTEFP:     "vrsqrtefp",
	VCIPHER:       "vcipher",
	VCIPHERLAST:   "vcipherlast",
	VNCIPHER:      "vncipher",
	VNCIPHERLAST:  "vncipherlast",
	VSBOX:         "vsbox",
	VSHASIGMAD:    "vshasigmad",
	VSHASIGMAW:    "vshasigmaw",
	VPMSUMB:       "vpmsumb",
	VPMSUMD:       "vpmsumd",
	VPMSUMH:       "vpmsumh",
	VPMSUMW:       "vpmsumw",
	VPERMXOR:      "vpermxor",
	VGBBD:         "vgbbd",
	VCLZB:         "vclzb",
	VCLZH:         "vclzh",
	VCLZW:         "vclzw",
	VCLZD:         "vclzd",
	VPOPCNTB:      "vpopcntb",
	VPOPCNTD:      "vpopcntd",
	VPOPCNTH:      "vpopcnth",
	VPOPCNTW:      "vpopcntw",
	VBPERMQ:       "vbpermq",
	BCDADD_:       "bcdadd.",
	BCDSUB_:       "bcdsub.",
	MTVSCR:        "mtvscr",
	MFVSCR:        "mfvscr",
	DADD:          "dadd",
	DADD_:         "dadd.",
	DSUB:          "dsub",
	DSUB_:         "dsub.",
	DMUL:          "dmul",
	DMUL_:         "dmul.",
	DDIV:          "ddiv",
	DDIV_:         "ddiv.",
	DCMPU:         "dcmpu",
	DCMPO:         "dcmpo",
	DTSTDC:        "dtstdc",
	DTSTDG:        "dtstdg",
	DTSTEX:        "dtstex",
	DTSTSF:        "dtstsf",
	DQUAI:         "dquai",
	DQUAI_:        "dquai.",
	DQUA:          "dqua",
	DQUA_:         "dqua.",
	DRRND:         "drrnd",
	DRRND_:        "drrnd.",
	DRINTX:        "drintx",
	DRINTX_:       "drintx.",
	DRINTN:        "drintn",
	DRINTN_:       "drintn.",
	DCTDP:         "dctdp",
	DCTDP_:        "dctdp.",
	DCTQPQ:        "dctqpq",
	DCTQPQ_:       "dctqpq.",
	DRSP:          "drsp",
	DRSP_:         "drsp.",
	DRDPQ:         "drdpq",
	DRDPQ_:        "drdpq.",
	DCFFIX:        "dcffix",
	DCFFIX_:       "dcffix.",
	DCFFIXQ:       "dcffixq",
	DCFFIXQ_:      "dcffixq.",
	DCTFIX:        "dctfix",
	DCTFIX_:       "dctfix.",
	DDEDPD:        "ddedpd",
	DDEDPD_:       "ddedpd.",
	DENBCD:        "denbcd",
	DENBCD_:       "denbcd.",
	DXEX:          "dxex",
	DXEX_:         "dxex.",
	DIEX:          "diex",
	DIEX_:         "diex.",
	DSCLI:         "dscli",
	DSCLI_:        "dscli.",
	DSCRI:         "dscri",
	DSCRI_:        "dscri.",
	LXSDX:         "lxsdx",
	LXSIWAX:       "lxsiwax",
	LXSIWZX:       "lxsiwzx",
	LXSSPX:        "lxsspx",
	LXVD2X:        "lxvd2x",
	LXVDSX:        "lxvdsx",
	LXVW4X:        "lxvw4x",
	STXSDX:        "stxsdx",
	STXSIWX:       "stxsiwx",
	STXSSPX:       "stxsspx",
	STXVD2X:       "stxvd2x",
	STXVW4X:       "stxvw4x",
	XSABSDP:       "xsabsdp",
	XSADDDP:       "xsadddp",
	XSADDSP:       "xsaddsp",
	XSCMPODP:      "xscmpodp",
	XSCMPUDP:      "xscmpudp",
	XSCPSGNDP:     "xscpsgndp",
	XSCVDPSP:      "xscvdpsp",
	XSCVDPSPN:     "xscvdpspn",
	XSCVDPSXDS:    "xscvdpsxds",
	XSCVDPSXWS:    "xscvdpsxws",
	XSCVDPUXDS:    "xscvdpuxds",
	XSCVDPUXWS:    "xscvdpuxws",
	XSCVSPDP:      "xscvspdp",
	XSCVSPDPN:     "xscvspdpn",
	XSCVSXDDP:     "xscvsxddp",
	XSCVSXDSP:     "xscvsxdsp",
	XSCVUXDDP:     "xscvuxddp",
	XSCVUXDSP:     "xscvuxdsp",
	XSDIVDP:       "xsdivdp",
	XSDIVSP:       "xsdivsp",
	XSMADDADP:     "xsmaddadp",
	XSMADDASP:     "xsmaddasp",
	XSMAXDP:       "xsmaxdp",
	XSMINDP:       "xsmindp",
	XSMSUBADP:     "xsmsubadp",
	XSMSUBASP:     "xsmsubasp",
	XSMULDP:       "xsmuldp",
	XSMULSP:       "xsmulsp",
	XSNABSDP:      "xsnabsdp",
	XSNEGDP:       "xsnegdp",
	XSNMADDADP:    "xsnmaddadp",
	XSNMADDASP:    "xsnmaddasp",
	XSNMSUBADP:    "xsnmsubadp",
	XSNMSUBASP:    "xsnmsubasp",
	XSRDPI:        "xsrdpi",
	XSRDPIC:       "xsrdpic",
	XSRDPIM:       "xsrdpim",
	XSRDPIP:       "xsrdpip",
	XSRDPIZ:       "xsrdpiz",
	XSREDP:        "xsredp",
	XSRESP:        "xsresp",
	XSRSP:         "xsrsp",
	XSRSQRTEDP:    "xsrsqrtedp",
	XSRSQRTESP:    "xsrsqrtesp",
	XSSQRTDP:      "xssqrtdp",
	XSSQRTSP:      "xssqrtsp",
	XSSUBDP:       "xssubdp",
	XSSUBSP:       "xssubsp",
	XSTDIVDP:      "xstdivdp",
	XSTSQRTDP:     "xstsqrtdp",
	XVABSDP:       "xvabsdp",
	XVABSSP:       "xvabssp",
	XVADDDP:       "xvadddp",
	XVADDSP:       "xvaddsp",
	XVCMPEQDP:     "xvcmpeqdp",
	XVCMPEQDP_:    "xvcmpeqdp.",
	XVCMPEQSP:     "xvcmpeqsp",
	XVCMPEQSP_:    "xvcmpeqsp.",
	XVCMPGEDP:     "xvcmpgedp",
	XVCMPGEDP_:    "xvcmpgedp.",
	XVCMPGESP:     "xvcmpgesp",
	XVCMPGESP_:    "xvcmpgesp.",
	XVCMPGTDP:     "xvcmpgtdp",
	XVCMPGTDP_:    "xvcmpgtdp.",
	XVCMPGTSP:     "xvcmpgtsp",
	XVCMPGTSP_:    "xvcmpgtsp.",
	XVCPSGNDP:     "xvcpsgndp",
	XVCPSGNSP:     "xvcpsgnsp",
	XVCVDPSP:      "xvcvdpsp",
	XVCVDPSXDS:    "xvcvdpsxds",
	XVCVDPSXWS:    "xvcvdpsxws",
	XVCVDPUXDS:    "xvcvdpuxds",
	XVCVDPUXWS:    "xvcvdpuxws",
	XVCVSPDP:      "xvcvspdp",
	XVCVSPSXDS:    "xvcvspsxds",
	XVCVSPSXWS:    "xvcvspsxws",
	XVCVSPUXDS:    "xvcvspuxds",
	XVCVSPUXWS:    "xvcvspuxws",
	XVCVSXDDP:     "xvcvsxddp",
	XVCVSXDSP:     "xvcvsxdsp",
	XVCVSXWDP:     "xvcvsxwdp",
	XVCVSXWSP:     "xvcvsxwsp",
	XVCVUXDDP:     "xvcvuxddp",
	XVCVUXDSP:     "xvcvuxdsp",
	XVCVUXWDP:     "xvcvuxwdp",
	XVCVUXWSP:     "xvcvuxwsp",
	XVDIVDP:       "xvdivdp",
	XVDIVSP:       "xvdivsp",
	XVMADDADP:     "xvmaddadp",
	XVMADDASP:     "xvmaddasp",
	XVMAXDP:       "xvmaxdp",
	XVMAXSP:       "xvmaxsp",
	XVMINDP:       "xvmindp",
	XVMINSP:       "xvminsp",
	XVMSUBADP:     "xvmsubadp",
	XVMSUBASP:     "xvmsubasp",
	XVMULDP:       "xvmuldp",
	XVMULSP:       "xvmulsp",
	XVNABSDP:      "xvnabsdp",
	XVNABSSP:      "xvnabssp",
	XVNEGDP:       "xvnegdp",
	XVNEGSP:       "xvnegsp",
	XVNMADDADP:    "xvnmaddadp",
	XVNMADDASP:    "xvnmaddasp",
	XVNMSUBADP:    "xvnmsubadp",
	XVNMSUBASP:    "xvnmsubasp",
	XVRDPI:        "xvrdpi",
	XVRDPIC:       "xvrdpic",
	XVRDPIM:       "xvrdpim",
	XVRDPIP:       "xvrdpip",
	XVRDPIZ:       "xvrdpiz",
	XVREDP:        "xvredp",
	XVRESP:        "xvresp",
	XVRSPI:        "xvrspi",
	XVRSPIC:       "xvrspic",
	XVRSPIM:       "xvrspim",
	XVRSPIP:       "xvrspip",
	XVRSPIZ:       "xvrspiz",
	XVRSQRTEDP:    "xvrsqrtedp",
	XVRSQRTESP:    "xvrsqrtesp",
	XVSQRTDP:      "xvsqrtdp",
	XVSQRTSP:      "xvsqrtsp",
	XVSUBDP:       "xvsubdp",
	XVSUBSP:       "xvsubsp",
	XVTDIVDP:      "xvtdivdp",
	XVTDIVSP:      "xvtdivsp",
	XVTSQRTDP:     "xvtsqrtdp",
	XVTSQRTSP:     "xvtsqrtsp",
	XXLAND:        "xxland",
	XXLANDC:       "xxlandc",
	XXLEQV:        "xxleqv",
	XXLNAND:       "xxlnand",
	XXLORC:        "xxlorc",
	XXLNOR:        "xxlnor",
	XXLOR:         "xxlor",
	XXLXOR:        "xxlxor",
	XXMRGHW:       "xxmrghw",
	XXMRGLW:       "xxmrglw",
	XXPERMDI:      "xxpermdi",
	XXSEL:         "xxsel",
	XXSLDWI:       "xxsldwi",
	XXSPLTW:       "xxspltw",
	BRINC:         "brinc",
	EVABS:         "evabs",
	EVADDIW:       "evaddiw",
	EVADDSMIAAW:   "evaddsmiaaw",
	EVADDSSIAAW:   "evaddssiaaw",
	EVADDUMIAAW:   "evaddumiaaw",
	EVADDUSIAAW:   "evaddusiaaw",
	EVADDW:        "evaddw",
	EVAND:         "evand",
	EVCMPEQ:       "evcmpeq",
	EVANDC:        "evandc",
	EVCMPGTS:      "evcmpgts",
	EVCMPGTU:      "evcmpgtu",
	EVCMPLTU:      "evcmpltu",
	EVCMPLTS:      "evcmplts",
	EVCNTLSW:      "evcntlsw",
	EVCNTLZW:      "evcntlzw",
	EVDIVWS:       "evdivws",
	EVDIVWU:       "evdivwu",
	EVEQV:         "eveqv",
	EVEXTSB:       "evextsb",
	EVEXTSH:       "evextsh",
	EVLDD:         "evldd",
	EVLDH:         "evldh",
	EVLDDX:        "evlddx",
	EVLDHX:        "evldhx",
	EVLDW:         "evldw",
	EVLHHESPLAT:   "evlhhesplat",
	EVLDWX:        "evldwx",
	EVLHHESPLATX:  "evlhhesplatx",
	EVLHHOSSPLAT:  "evlhhossplat",
	EVLHHOUSPLAT:  "evlhhousplat",
	EVLHHOSSPLATX: "evlhhossplatx",
	EVLHHOUSPLATX: "evlhhousplatx",
	EVLWHE:        "evlwhe",
	EVLWHOS:       "evlwhos",
	EVLWHEX:       "evlwhex",
	EVLWHOSX:      "evlwhosx",
	EVLWHOU:       "evlwhou",
	EVLWHSPLAT:    "evlwhsplat",
	EVLWHOUX:      "evlwhoux",
	EVLWHSPLATX:   "evlwhsplatx",
	EVLWWSPLAT:    "evlwwsplat",
	EVMERGEHI:     "evmergehi",
	EVLWWSPLATX:   "evlwwsplatx",
	EVMERGELO:     "evmergelo",
	EVMERGEHILO:   "evmergehilo",
	EVMHEGSMFAA:   "evmhegsmfaa",
	EVMERGELOHI:   "evmergelohi",
	EVMHEGSMFAN:   "evmhegsmfan",
	EVMHEGSMIAA:   "evmhegsmiaa",
	EVMHEGUMIAA:   "evmhegumiaa",
	EVMHEGSMIAN:   "evmhegsmian",
	EVMHEGUMIAN:   "evmhegumian",
	EVMHESMF:      "evmhesmf",
	EVMHESMFAAW:   "evmhesmfaaw",
	EVMHESMFA:     "evmhesmfa",
	EVMHESMFANW:   "evmhesmfanw",
	EVMHESMI:      "evmhesmi",
	EVMHESMIAAW:   "evmhesmiaaw",
	EVMHESMIA:     "evmhesmia",
	EVMHESMIANW:   "evmhesmianw",
	EVMHESSF:      "evmhessf",
	EVMHESSFA:     "evmhessfa",
	EVMHESSFAAW:   "evmhessfaaw",
	EVMHESSFANW:   "evmhessfanw",
	EVMHESSIAAW:   "evmhessiaaw",
	EVMHESSIANW:   "evmhessianw",
	EVMHEUMI:      "evmheumi",
	EVMHEUMIAAW:   "evmheumiaaw",
	EVMHEUMIA:     "evmheumia",
	EVMHEUMIANW:   "evmheumianw",
	EVMHEUSIAAW:   "evmheusiaaw",
	EVMHEUSIANW:   "evmheusianw",
	EVMHOGSMFAA:   "evmhogsmfaa",
	EVMHOGSMIAA:   "evmhogsmiaa",
	EVMHOGSMFAN:   "evmhogsmfan",
	EVMHOGSMIAN:   "evmhogsmian",
	EVMHOGUMIAA:   "evmhogumiaa",
	EVMHOSMF:      "evmhosmf",
	EVMHOGUMIAN:   "evmhogumian",
	EVMHOSMFA:     "evmhosmfa",
	EVMHOSMFAAW:   "evmhosmfaaw",
	EVMHOSMI:      "evmhosmi",
	EVMHOSMFANW:   "evmhosmfanw",
	EVMHOSMIA:     "evmhosmia",
	EVMHOSMIAAW:   "evmhosmiaaw",
	EVMHOSMIANW:   "evmhosmianw",
	EVMHOSSF:      "evmhossf",
	EVMHOSSFA:     "evmhossfa",
	EVMHOSSFAAW:   "evmhossfaaw",
	EVMHOSSFANW:   "evmhossfanw",
	EVMHOSSIAAW:   "evmhossiaaw",
	EVMHOUMI:      "evmhoumi",
	EVMHOSSIANW:   "evmhossianw",
	EVMHOUMIA:     "evmhoumia",
	EVMHOUMIAAW:   "evmhoumiaaw",
	EVMHOUSIAAW:   "evmhousiaaw",
	EVMHOUMIANW:   "evmhoumianw",
	EVMHOUSIANW:   "evmhousianw",
	EVMRA:         "evmra",
	EVMWHSMF:      "evmwhsmf",
	EVMWHSMI:      "evmwhsmi",
	EVMWHSMFA:     "evmwhsmfa",
	EVMWHSMIA:     "evmwhsmia",
	EVMWHSSF:      "evmwhssf",
	EVMWHUMI:      "evmwhumi",
	EVMWHSSFA:     "evmwhssfa",
	EVMWHUMIA:     "evmwhumia",
	EVMWLSMIAAW:   "evmwlsmiaaw",
	EVMWLSSIAAW:   "evmwlssiaaw",
	EVMWLSMIANW:   "evmwlsmianw",
	EVMWLSSIANW:   "evmwlssianw",
	EVMWLUMI:      "evmwlumi",
	EVMWLUMIAAW:   "evmwlumiaaw",
	EVMWLUMIA:     "evmwlumia",
	EVMWLUMIANW:   "evmwlumianw",
	EVMWLUSIAAW:   "evmwlusiaaw",
	EVMWSMF:       "evmwsmf",
	EVMWLUSIANW:   "evmwlusianw",
	EVMWSMFA:      "evmwsmfa",
	EVMWSMFAA:     "evmwsmfaa",
	EVMWSMI:       "evmwsmi",
	EVMWSMIAA:     "evmwsmiaa",
	EVMWSMFAN:     "evmwsmfan",
	EVMWSMIA:      "evmwsmia",
	EVMWSMIAN:     "evmwsmian",
	EVMWSSF:       "evmwssf",
	EVMWSSFA:      "evmwssfa",
	EVMWSSFAA:     "evmwssfaa",
	EVMWUMI:       "evmwumi",
	EVMWSSFAN:     "evmwssfan",
	EVMWUMIA:      "evmwumia",
	EVMWUMIAA:     "evmwumiaa",
	EVNAND:        "evnand",
	EVMWUMIAN:     "evmwumian",
	EVNEG:         "evneg",
	EVNOR:         "evnor",
	EVORC:         "evorc",
	EVOR:          "evor",
	EVRLW:         "evrlw",
	EVRLWI:        "evrlwi",
	EVSEL:         "evsel",
	EVRNDW:        "evrndw",
	EVSLW:         "evslw",
	EVSPLATFI:     "evsplatfi",
	EVSRWIS:       "evsrwis",
	EVSLWI:        "evslwi",
	EVSPLATI:      "evsplati",
	EVSRWIU:       "evsrwiu",
	EVSRWS:        "evsrws",
	EVSTDD:        "evstdd",
	EVSRWU:        "evsrwu",
	EVSTDDX:       "evstddx",
	EVSTDH:        "evstdh",
	EVSTDW:        "evstdw",
	EVSTDHX:       "evstdhx",
	EVSTDWX:       "evstdwx",
	EVSTWHE:       "evstwhe",
	EVSTWHO:       "evstwho",
	EVSTWWE:       "evstwwe",
	EVSTWHEX:      "evstwhex",
	EVSTWHOX:      "evstwhox",
	EVSTWWEX:      "evstwwex",
	EVSTWWO:       "evstwwo",
	EVSUBFSMIAAW:  "evsubfsmiaaw",
	EVSTWWOX:      "evstwwox",
	EVSUBFSSIAAW:  "evsubfssiaaw",
	EVSUBFUMIAAW:  "evsubfumiaaw",
	EVSUBFUSIAAW:  "evsubfusiaaw",
	EVSUBFW:       "evsubfw",
	EVSUBIFW:      "evsubifw",
	EVXOR:         "evxor",
	EVFSABS:       "evfsabs",
	EVFSNABS:      "evfsnabs",
	EVFSNEG:       "evfsneg",
	EVFSADD:       "evfsadd",
	EVFSMUL:       "evfsmul",
	EVFSSUB:       "evfssub",
	EVFSDIV:       "evfsdiv",
	EVFSCMPGT:     "evfscmpgt",
	EVFSCMPLT:     "evfscmplt",
	EVFSCMPEQ:     "evfscmpeq",
	EVFSTSTGT:     "evfststgt",
	EVFSTSTLT:     "evfststlt",
	EVFSTSTEQ:     "evfststeq",
	EVFSCFSI:      "evfscfsi",
	EVFSCFSF:      "evfscfsf",
	EVFSCFUI:      "evfscfui",
	EVFSCFUF:      "evfscfuf",
	EVFSCTSI:      "evfsctsi",
	EVFSCTUI:      "evfsctui",
	EVFSCTSIZ:     "evfsctsiz",
	EVFSCTUIZ:     "evfsctuiz",
	EVFSCTSF:      "evfsctsf",
	EVFSCTUF:      "evfsctuf",
	EFSABS:        "efsabs",
	EFSNEG:        "efsneg",
	EFSNABS:       "efsnabs",
	EFSADD:        "efsadd",
	EFSMUL:        "efsmul",
	EFSSUB:        "efssub",
	EFSDIV:        "efsdiv",
	EFSCMPGT:      "efscmpgt",
	EFSCMPLT:      "efscmplt",
	EFSCMPEQ:      "efscmpeq",
	EFSTSTGT:      "efststgt",
	EFSTSTLT:      "efststlt",
	EFSTSTEQ:      "efststeq",
	EFSCFSI:       "efscfsi",
	EFSCFSF:       "efscfsf",
	EFSCTSI:       "efsctsi",
	EFSCFUI:       "efscfui",
	EFSCFUF:       "efscfuf",
	EFSCTUI:       "efsctui",
	EFSCTSIZ:      "efsctsiz",
	EFSCTSF:       "efsctsf",
	EFSCTUIZ:      "efsctuiz",
	EFSCTUF:       "efsctuf",
	EFDABS:        "efdabs",
	EFDNEG:        "efdneg",
	EFDNABS:       "efdnabs",
	EFDADD:        "efdadd",
	EFDMUL:        "efdmul",
	EFDSUB:        "efdsub",
	EFDDIV:        "efddiv",
	EFDCMPGT:      "efdcmpgt",
	EFDCMPEQ:      "efdcmpeq",
	EFDCMPLT:      "efdcmplt",
	EFDTSTGT:      "efdtstgt",
	EFDTSTLT:      "efdtstlt",
	EFDCFSI:       "efdcfsi",
	EFDTSTEQ:      "efdtsteq",
	EFDCFUI:       "efdcfui",
	EFDCFSID:      "efdcfsid",
	EFDCFSF:       "efdcfsf",
	EFDCFUF:       "efdcfuf",
	EFDCFUID:      "efdcfuid",
	EFDCTSI:       "efdctsi",
	EFDCTUI:       "efdctui",
	EFDCTSIDZ:     "efdctsidz",
	EFDCTUIDZ:     "efdctuidz",
	EFDCTSIZ:      "efdctsiz",
	EFDCTSF:       "efdctsf",
	EFDCTUF:       "efdctuf",
	EFDCTUIZ:      "efdctuiz",
	EFDCFS:        "efdcfs",
	EFSCFD:        "efscfd",
	DLMZB:         "dlmzb",
	DLMZB_:        "dlmzb.",
	MACCHW:        "macchw",
	MACCHW_:       "macchw.",
	MACCHWO:       "macchwo",
	MACCHWO_:      "macchwo.",
	MACCHWS:       "macchws",
	MACCHWS_:      "macchws.",
	MACCHWSO:      "macchwso",
	MACCHWSO_:     "macchwso.",
	MACCHWU:       "macchwu",
	MACCHWU_:      "macchwu.",
	MACCHWUO:      "macchwuo",
	MACCHWUO_:     "macchwuo.",
	MACCHWSU:      "macchwsu",
	MACCHWSU_:     "macchwsu.",
	MACCHWSUO:     "macchwsuo",
	MACCHWSUO_:    "macchwsuo.",
	MACHHW:        "machhw",
	MACHHW_:       "machhw.",
	MACHHWO:       "machhwo",
	MACHHWO_:      "machhwo.",
	MACHHWS:       "machhws",
	MACHHWS_:      "machhws.",
	MACHHWSO:      "machhwso",
	MACHHWSO_:     "machhwso.",
	MACHHWU:       "machhwu",
	MACHHWU_:      "machhwu.",
	MACHHWUO:      "machhwuo",
	MACHHWUO_:     "machhwuo.",
	MACHHWSU:      "machhwsu",
	MACHHWSU_:     "machhwsu.",
	MACHHWSUO:     "machhwsuo",
	MACHHWSUO_:    "machhwsuo.",
	MACLHW:        "maclhw",
	MACLHW_:       "maclhw.",
	MACLHWO:       "maclhwo",
	MACLHWO_:      "maclhwo.",
	MACLHWS:       "maclhws",
	MACLHWS_:      "maclhws.",
	MACLHWSO:      "maclhwso",
	MACLHWSO_:     "maclhwso.",
	MACLHWU:       "maclhwu",
	MACLHWU_:      "maclhwu.",
	MACLHWUO:      "maclhwuo",
	MACLHWUO_:     "maclhwuo.",
	MULCHW:        "mulchw",
	MULCHW_:       "mulchw.",
	MACLHWSU:      "maclhwsu",
	MACLHWSU_:     "maclhwsu.",
	MACLHWSUO:     "maclhwsuo",
	MACLHWSUO_:    "maclhwsuo.",
	MULCHWU:       "mulchwu",
	MULCHWU_:      "mulchwu.",
	MULHHW:        "mulhhw",
	MULHHW_:       "mulhhw.",
	MULLHW:        "mullhw",
	MULLHW_:       "mullhw.",
	MULHHWU:       "mulhhwu",
	MULHHWU_:      "mulhhwu.",
	MULLHWU:       "mullhwu",
	MULLHWU_:      "mullhwu.",
	NMACCHW:       "nmacchw",
	NMACCHW_:      "nmacchw.",
	NMACCHWO:      "nmacchwo",
	NMACCHWO_:     "nmacchwo.",
	NMACCHWS:      "nmacchws",
	NMACCHWS_:     "nmacchws.",
	NMACCHWSO:     "nmacchwso",
	NMACCHWSO_:    "nmacchwso.",
	NMACHHW:       "nmachhw",
	NMACHHW_:      "nmachhw.",
	NMACHHWO:      "nmachhwo",
	NMACHHWO_:     "nmachhwo.",
	NMACHHWS:      "nmachhws",
	NMACHHWS_:     "nmachhws.",
	NMACHHWSO:     "nmachhwso",
	NMACHHWSO_:    "nmachhwso.",
	NMACLHW:       "nmaclhw",
	NMACLHW_:      "nmaclhw.",
	NMACLHWO:      "nmaclhwo",
	NMACLHWO_:     "nmaclhwo.",
	NMACLHWS:      "nmaclhws",
	NMACLHWS_:     "nmaclhws.",
	NMACLHWSO:     "nmaclhwso",
	NMACLHWSO_:    "nmaclhwso.",
	ICBI:          "icbi",
	ICBT:          "icbt",
	DCBA:          "dcba",
	DCBT:          "dcbt",
	DCBTST:        "dcbtst",
	DCBZ:          "dcbz",
	DCBST:         "dcbst",
	DCBF:          "dcbf",
	ISYNC:         "isync",
	LBARX:         "lbarx",
	LHARX:         "lharx",
	LWARX:         "lwarx",
	STBCX_:        "stbcx.",
	STHCX_:        "sthcx.",
	STWCX_:        "stwcx.",
	LDARX:         "ldarx",
	STDCX_:        "stdcx.",
	LQARX:         "lqarx",
	STQCX_:        "stqcx.",
	SYNC:          "sync",
	EIEIO:         "eieio",
	MBAR:          "mbar",
	WAIT:          "wait",
	TBEGIN_:       "tbegin.",
	TEND_:         "tend.",
	TABORT_:       "tabort.",
	TABORTWC_:     "tabortwc.",
	TABORTWCI_:    "tabortwci.",
	TABORTDC_:     "tabortdc.",
	TABORTDCI_:    "tabortdci.",
	TSR_:          "tsr.",
	TCHECK:        "tcheck",
	MFTB:          "mftb",
	RFEBB:         "rfebb",
	LBDX:          "lbdx",
	LHDX:          "lhdx",
	LWDX:          "lwdx",
	LDDX:          "lddx",
	LFDDX:         "lfddx",
	STBDX:         "stbdx",
	STHDX:         "sthdx",
	STWDX:         "stwdx",
	STDDX:         "stddx",
	STFDDX:        "stfddx",
	DSN:           "dsn",
	ECIWX:         "eciwx",
	ECOWX:         "ecowx",
	RFID:          "rfid",
	HRFID:         "hrfid",
	DOZE:          "doze",
	NAP:           "nap",
	SLEEP:         "sleep",
	RVWINKLE:      "rvwinkle",
	LBZCIX:        "lbzcix",
	LWZCIX:        "lwzcix",
	LHZCIX:        "lhzcix",
	LDCIX:         "ldcix",
	STBCIX:        "stbcix",
	STWCIX:        "stwcix",
	STHCIX:        "sthcix",
	STDCIX:        "stdcix",
	TRECLAIM_:     "treclaim.",
	TRECHKPT_:     "trechkpt.",
	MTMSR:         "mtmsr",
	MTMSRD:        "mtmsrd",
	MFMSR:         "mfmsr",
	SLBIE:         "slbie",
	SLBIA:         "slbia",
	SLBMTE:        "slbmte",
	SLBMFEV:       "slbmfev",
	SLBMFEE:       "slbmfee",
	SLBFEE_:       "slbfee.",
	MTSR:          "mtsr",
	MTSRIN:        "mtsrin",
	MFSR:          "mfsr",
	MFSRIN:        "mfsrin",
	TLBIE:         "tlbie",
	TLBIEL:        "tlbiel",
	TLBIA:         "tlbia",
	TLBSYNC:       "tlbsync",
	MSGSND:        "msgsnd",
	MSGCLR:        "msgclr",
	MSGSNDP:       "msgsndp",
	MSGCLRP:       "msgclrp",
	MTTMR:         "mttmr",
	RFI:           "rfi",
	RFCI:          "rfci",
	RFDI:          "rfdi",
	RFMCI:         "rfmci",
	RFGI:          "rfgi",
	EHPRIV:        "ehpriv",
	MTDCR:         "mtdcr",
	MTDCRX:        "mtdcrx",
	MFDCR:         "mfdcr",
	MFDCRX:        "mfdcrx",
	WRTEE:         "wrtee",
	WRTEEI:        "wrteei",
	LBEPX:         "lbepx",
	LHEPX:         "lhepx",
	LWEPX:         "lwepx",
	LDEPX:         "ldepx",
	STBEPX:        "stbepx",
	STHEPX:        "sthepx",
	STWEPX:        "stwepx",
	STDEPX:        "stdepx",
	DCBSTEP:       "dcbstep",
	DCBTEP:        "dcbtep",
	DCBFEP:        "dcbfep",
	DCBTSTEP:      "dcbtstep",
	ICBIEP:        "icbiep",
	DCBZEP:        "dcbzep",
	LFDEPX:        "lfdepx",
	STFDEPX:       "stfdepx",
	EVLDDEPX:      "evlddepx",
	EVSTDDEPX:     "evstddepx",
	LVEPX:         "lvepx",
	LVEPXL:        "lvepxl",
	STVEPX:        "stvepx",
	STVEPXL:       "stvepxl",
	DCBI:          "dcbi",
	DCBLQ_:        "dcblq.",
	ICBLQ_:        "icblq.",
	DCBTLS:        "dcbtls",
	DCBTSTLS:      "dcbtstls",
	ICBTLS:        "icbtls",
	ICBLC:         "icblc",
	DCBLC:         "dcblc",
	TLBIVAX:       "tlbivax",
	TLBILX:        "tlbilx",
	TLBSX:         "tlbsx",
	TLBSRX_:       "tlbsrx.",
	TLBRE:         "tlbre",
	TLBWE:         "tlbwe",
	DNH:           "dnh",
	DCI:           "dci",
	ICI:           "ici",
	DCREAD:        "dcread",
	ICREAD:        "icread",
	MFPMR:         "mfpmr",
	MTPMR:         "mtpmr",
}

var (
	ap_Reg_11_15               = &argField{Type: TypeReg, Shift: 0, BitFields: BitFields{{11, 5}}}
	ap_Reg_6_10                = &argField{Type: TypeReg, Shift: 0, BitFields: BitFields{{6, 5}}}
	ap_PCRel_6_29_shift2       = &argField{Type: TypePCRel, Shift: 2, BitFields: BitFields{{6, 24}}}
	ap_Label_6_29_shift2       = &argField{Type: TypeLabel, Shift: 2, BitFields: BitFields{{6, 24}}}
	ap_ImmUnsigned_6_10        = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{6, 5}}}
	ap_CondRegBit_11_15        = &argField{Type: TypeCondRegBit, Shift: 0, BitFields: BitFields{{11, 5}}}
	ap_PCRel_16_29_shift2      = &argField{Type: TypePCRel, Shift: 2, BitFields: BitFields{{16, 14}}}
	ap_Label_16_29_shift2      = &argField{Type: TypeLabel, Shift: 2, BitFields: BitFields{{16, 14}}}
	ap_ImmUnsigned_19_20       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{19, 2}}}
	ap_CondRegBit_6_10         = &argField{Type: TypeCondRegBit, Shift: 0, BitFields: BitFields{{6, 5}}}
	ap_CondRegBit_16_20        = &argField{Type: TypeCondRegBit, Shift: 0, BitFields: BitFields{{16, 5}}}
	ap_CondRegField_6_8        = &argField{Type: TypeCondRegField, Shift: 0, BitFields: BitFields{{6, 3}}}
	ap_CondRegField_11_13      = &argField{Type: TypeCondRegField, Shift: 0, BitFields: BitFields{{11, 3}}}
	ap_ImmUnsigned_20_26       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{20, 7}}}
	ap_SpReg_11_20             = &argField{Type: TypeSpReg, Shift: 0, BitFields: BitFields{{11, 10}}}
	ap_Offset_16_31            = &argField{Type: TypeOffset, Shift: 0, BitFields: BitFields{{16, 16}}}
	ap_Reg_16_20               = &argField{Type: TypeReg, Shift: 0, BitFields: BitFields{{16, 5}}}
	ap_Offset_16_29_shift2     = &argField{Type: TypeOffset, Shift: 2, BitFields: BitFields{{16, 14}}}
	ap_Offset_16_27_shift4     = &argField{Type: TypeOffset, Shift: 4, BitFields: BitFields{{16, 12}}}
	ap_ImmUnsigned_16_20       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{16, 5}}}
	ap_ImmSigned_16_31         = &argField{Type: TypeImmSigned, Shift: 0, BitFields: BitFields{{16, 16}}}
	ap_ImmUnsigned_16_31       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{16, 16}}}
	ap_CondRegBit_21_25        = &argField{Type: TypeCondRegBit, Shift: 0, BitFields: BitFields{{21, 5}}}
	ap_ImmUnsigned_21_25       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{21, 5}}}
	ap_ImmUnsigned_26_30       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{26, 5}}}
	ap_ImmUnsigned_30_30_16_20 = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{30, 1}, {16, 5}}}
	ap_ImmUnsigned_26_26_21_25 = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{26, 1}, {21, 5}}}
	ap_SpReg_16_20_11_15       = &argField{Type: TypeSpReg, Shift: 0, BitFields: BitFields{{16, 5}, {11, 5}}}
	ap_ImmUnsigned_12_19       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{12, 8}}}
	ap_ImmUnsigned_10_10       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{10, 1}}}
	ap_VecReg_31_31_6_10       = &argField{Type: TypeVecReg, Shift: 0, BitFields: BitFields{{31, 1}, {6, 5}}}
	ap_FPReg_6_10              = &argField{Type: TypeFPReg, Shift: 0, BitFields: BitFields{{6, 5}}}
	ap_FPReg_16_20             = &argField{Type: TypeFPReg, Shift: 0, BitFields: BitFields{{16, 5}}}
	ap_FPReg_11_15             = &argField{Type: TypeFPReg, Shift: 0, BitFields: BitFields{{11, 5}}}
	ap_FPReg_21_25             = &argField{Type: TypeFPReg, Shift: 0, BitFields: BitFields{{21, 5}}}
	ap_ImmUnsigned_16_19       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{16, 4}}}
	ap_ImmUnsigned_15_15       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{15, 1}}}
	ap_ImmUnsigned_7_14        = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{7, 8}}}
	ap_ImmUnsigned_6_6         = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{6, 1}}}
	ap_VecReg_6_10             = &argField{Type: TypeVecReg, Shift: 0, BitFields: BitFields{{6, 5}}}
	ap_VecReg_11_15            = &argField{Type: TypeVecReg, Shift: 0, BitFields: BitFields{{11, 5}}}
	ap_VecReg_16_20            = &argField{Type: TypeVecReg, Shift: 0, BitFields: BitFields{{16, 5}}}
	ap_ImmUnsigned_12_15       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{12, 4}}}
	ap_ImmUnsigned_13_15       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{13, 3}}}
	ap_ImmUnsigned_14_15       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{14, 2}}}
	ap_ImmSigned_11_15         = &argField{Type: TypeImmSigned, Shift: 0, BitFields: BitFields{{11, 5}}}
	ap_VecReg_21_25            = &argField{Type: TypeVecReg, Shift: 0, BitFields: BitFields{{21, 5}}}
	ap_ImmUnsigned_22_25       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{22, 4}}}
	ap_ImmUnsigned_11_15       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{11, 5}}}
	ap_ImmUnsigned_16_16       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{16, 1}}}
	ap_ImmUnsigned_17_20       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{17, 4}}}
	ap_ImmUnsigned_22_22       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{22, 1}}}
	ap_ImmUnsigned_16_21       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{16, 6}}}
	ap_ImmUnsigned_21_22       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{21, 2}}}
	ap_ImmUnsigned_11_12       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{11, 2}}}
	ap_ImmUnsigned_11_11       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{11, 1}}}
	ap_VecReg_30_30_16_20      = &argField{Type: TypeVecReg, Shift: 0, BitFields: BitFields{{30, 1}, {16, 5}}}
	ap_VecReg_29_29_11_15      = &argField{Type: TypeVecReg, Shift: 0, BitFields: BitFields{{29, 1}, {11, 5}}}
	ap_ImmUnsigned_22_23       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{22, 2}}}
	ap_VecReg_28_28_21_25      = &argField{Type: TypeVecReg, Shift: 0, BitFields: BitFields{{28, 1}, {21, 5}}}
	ap_CondRegField_29_31      = &argField{Type: TypeCondRegField, Shift: 0, BitFields: BitFields{{29, 3}}}
	ap_ImmUnsigned_7_10        = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{7, 4}}}
	ap_ImmUnsigned_9_10        = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{9, 2}}}
	ap_ImmUnsigned_31_31       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{31, 1}}}
	ap_ImmSigned_16_20         = &argField{Type: TypeImmSigned, Shift: 0, BitFields: BitFields{{16, 5}}}
	ap_ImmUnsigned_20_20       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{20, 1}}}
	ap_ImmUnsigned_8_10        = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{8, 3}}}
	ap_SpReg_12_15             = &argField{Type: TypeSpReg, Shift: 0, BitFields: BitFields{{12, 4}}}
	ap_ImmUnsigned_6_20        = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{6, 15}}}
	ap_ImmUnsigned_11_20       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{11, 10}}}
)

var instFormats = [...]instFormat{
	{CNTLZW, 0xfc0007ff, 0x7c000034, 0xf800, // Count Leading Zeros Word X-form (cntlzw RA, RS)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10}},
	{CNTLZW_, 0xfc0007ff, 0x7c000035, 0xf800, // Count Leading Zeros Word X-form (cntlzw. RA, RS)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10}},
	{B, 0xfc000003, 0x48000000, 0x0, // Branch I-form (b target_addr)
		[5]*argField{ap_PCRel_6_29_shift2}},
	{BA, 0xfc000003, 0x48000002, 0x0, // Branch I-form (ba target_addr)
		[5]*argField{ap_Label_6_29_shift2}},
	{BL, 0xfc000003, 0x48000001, 0x0, // Branch I-form (bl target_addr)
		[5]*argField{ap_PCRel_6_29_shift2}},
	{BLA, 0xfc000003, 0x48000003, 0x0, // Branch I-form (bla target_addr)
		[5]*argField{ap_Label_6_29_shift2}},
	{BC, 0xfc000003, 0x40000000, 0x0, // Branch Conditional B-form (bc BO,BI,target_addr)
		[5]*argField{ap_ImmUnsigned_6_10, ap_CondRegBit_11_15, ap_PCRel_16_29_shift2}},
	{BCA, 0xfc000003, 0x40000002, 0x0, // Branch Conditional B-form (bca BO,BI,target_addr)
		[5]*argField{ap_ImmUnsigned_6_10, ap_CondRegBit_11_15, ap_Label_16_29_shift2}},
	{BCL, 0xfc000003, 0x40000001, 0x0, // Branch Conditional B-form (bcl BO,BI,target_addr)
		[5]*argField{ap_ImmUnsigned_6_10, ap_CondRegBit_11_15, ap_PCRel_16_29_shift2}},
	{BCLA, 0xfc000003, 0x40000003, 0x0, // Branch Conditional B-form (bcla BO,BI,target_addr)
		[5]*argField{ap_ImmUnsigned_6_10, ap_CondRegBit_11_15, ap_Label_16_29_shift2}},
	{BCLR, 0xfc0007ff, 0x4c000020, 0xe000, // Branch Conditional to Link Register XL-form (bclr BO,BI,BH)
		[5]*argField{ap_ImmUnsigned_6_10, ap_CondRegBit_11_15, ap_ImmUnsigned_19_20}},
	{BCLRL, 0xfc0007ff, 0x4c000021, 0xe000, // Branch Conditional to Link Register XL-form (bclrl BO,BI,BH)
		[5]*argField{ap_ImmUnsigned_6_10, ap_CondRegBit_11_15, ap_ImmUnsigned_19_20}},
	{BCCTR, 0xfc0007ff, 0x4c000420, 0xe000, // Branch Conditional to Count Register XL-form (bcctr BO,BI,BH)
		[5]*argField{ap_ImmUnsigned_6_10, ap_CondRegBit_11_15, ap_ImmUnsigned_19_20}},
	{BCCTRL, 0xfc0007ff, 0x4c000421, 0xe000, // Branch Conditional to Count Register XL-form (bcctrl BO,BI,BH)
		[5]*argField{ap_ImmUnsigned_6_10, ap_CondRegBit_11_15, ap_ImmUnsigned_19_20}},
	{BCTAR, 0xfc0007ff, 0x4c000460, 0xe000, // Branch Conditional to Branch Target Address Register XL-form (bctar BO,BI,BH)
		[5]*argField{ap_ImmUnsigned_6_10, ap_CondRegBit_11_15, ap_ImmUnsigned_19_20}},
	{BCTARL, 0xfc0007ff, 0x4c000461, 0xe000, // Branch Conditional to Branch Target Address Register XL-form (bctarl BO,BI,BH)
		[5]*argField{ap_ImmUnsigned_6_10, ap_CondRegBit_11_15, ap_ImmUnsigned_19_20}},
	{CRAND, 0xfc0007fe, 0x4c000202, 0x1, // Condition Register AND XL-form (crand BT,BA,BB)
		[5]*argField{ap_CondRegBit_6_10, ap_CondRegBit_11_15, ap_CondRegBit_16_20}},
	{CROR, 0xfc0007fe, 0x4c000382, 0x1, // Condition Register OR XL-form (cror BT,BA,BB)
		[5]*argField{ap_CondRegBit_6_10, ap_CondRegBit_11_15, ap_CondRegBit_16_20}},
	{CRNAND, 0xfc0007fe, 0x4c0001c2, 0x1, // Condition Register NAND XL-form (crnand BT,BA,BB)
		[5]*argField{ap_CondRegBit_6_10, ap_CondRegBit_11_15, ap_CondRegBit_16_20}},
	{CRXOR, 0xfc0007fe, 0x4c000182, 0x1, // Condition Register XOR XL-form (crxor BT,BA,BB)
		[5]*argField{ap_CondRegBit_6_10, ap_CondRegBit_11_15, ap_CondRegBit_16_20}},
	{CRNOR, 0xfc0007fe, 0x4c000042, 0x1, // Condition Register NOR XL-form (crnor BT,BA,BB)
		[5]*argField{ap_CondRegBit_6_10, ap_CondRegBit_11_15, ap_CondRegBit_16_20}},
	{CRANDC, 0xfc0007fe, 0x4c000102, 0x1, // Condition Register AND with Complement XL-form (crandc BT,BA,BB)
		[5]*argField{ap_CondRegBit_6_10, ap_CondRegBit_11_15, ap_CondRegBit_16_20}},
	{MCRF, 0xfc0007fe, 0x4c000000, 0x63f801, // Move Condition Register Field XL-form (mcrf BF,BFA)
		[5]*argField{ap_CondRegField_6_8, ap_CondRegField_11_13}},
	{CREQV, 0xfc0007fe, 0x4c000242, 0x1, // Condition Register Equivalent XL-form (creqv BT,BA,BB)
		[5]*argField{ap_CondRegBit_6_10, ap_CondRegBit_11_15, ap_CondRegBit_16_20}},
	{CRORC, 0xfc0007fe, 0x4c000342, 0x1, // Condition Register OR with Complement XL-form (crorc BT,BA,BB)
		[5]*argField{ap_CondRegBit_6_10, ap_CondRegBit_11_15, ap_CondRegBit_16_20}},
	{SC, 0xfc000002, 0x44000002, 0x3fff01d, // System Call SC-form (sc LEV)
		[5]*argField{ap_ImmUnsigned_20_26}},
	{CLRBHRB, 0xfc0007fe, 0x7c00035c, 0x3fff801, // Clear BHRB X-form (clrbhrb)
		[5]*argField{}},
	{MFBHRBE, 0xfc0007fe, 0x7c00025c, 0x1, // Move From Branch History Rolling Buffer XFX-form (mfbhrbe RT,BHRBE)
		[5]*argField{ap_Reg_6_10, ap_SpReg_11_20}},
	{LBZ, 0xfc000000, 0x88000000, 0x0, // Load Byte and Zero D-form (lbz RT,D(RA))
		[5]*argField{ap_Reg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{LBZU, 0xfc000000, 0x8c000000, 0x0, // Load Byte and Zero with Update D-form (lbzu RT,D(RA))
		[5]*argField{ap_Reg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{LBZX, 0xfc0007fe, 0x7c0000ae, 0x1, // Load Byte and Zero Indexed X-form (lbzx RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LBZUX, 0xfc0007fe, 0x7c0000ee, 0x1, // Load Byte and Zero with Update Indexed X-form (lbzux RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LHZ, 0xfc000000, 0xa0000000, 0x0, // Load Halfword and Zero D-form (lhz RT,D(RA))
		[5]*argField{ap_Reg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{LHZU, 0xfc000000, 0xa4000000, 0x0, // Load Halfword and Zero with Update D-form (lhzu RT,D(RA))
		[5]*argField{ap_Reg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{LHZX, 0xfc0007fe, 0x7c00022e, 0x1, // Load Halfword and Zero Indexed X-form (lhzx RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LHZUX, 0xfc0007fe, 0x7c00026e, 0x1, // Load Halfword and Zero with Update Indexed X-form (lhzux RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LHA, 0xfc000000, 0xa8000000, 0x0, // Load Halfword Algebraic D-form (lha RT,D(RA))
		[5]*argField{ap_Reg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{LHAU, 0xfc000000, 0xac000000, 0x0, // Load Halfword Algebraic with Update D-form (lhau RT,D(RA))
		[5]*argField{ap_Reg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{LHAX, 0xfc0007fe, 0x7c0002ae, 0x1, // Load Halfword Algebraic Indexed X-form (lhax RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LHAUX, 0xfc0007fe, 0x7c0002ee, 0x1, // Load Halfword Algebraic with Update Indexed X-form (lhaux RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LWZ, 0xfc000000, 0x80000000, 0x0, // Load Word and Zero D-form (lwz RT,D(RA))
		[5]*argField{ap_Reg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{LWZU, 0xfc000000, 0x84000000, 0x0, // Load Word and Zero with Update D-form (lwzu RT,D(RA))
		[5]*argField{ap_Reg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{LWZX, 0xfc0007fe, 0x7c00002e, 0x1, // Load Word and Zero Indexed X-form (lwzx RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LWZUX, 0xfc0007fe, 0x7c00006e, 0x1, // Load Word and Zero with Update Indexed X-form (lwzux RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LWA, 0xfc000003, 0xe8000002, 0x0, // Load Word Algebraic DS-form (lwa RT,DS(RA))
		[5]*argField{ap_Reg_6_10, ap_Offset_16_29_shift2, ap_Reg_11_15}},
	{LWAX, 0xfc0007fe, 0x7c0002aa, 0x1, // Load Word Algebraic Indexed X-form (lwax RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LWAUX, 0xfc0007fe, 0x7c0002ea, 0x1, // Load Word Algebraic with Update Indexed X-form (lwaux RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LD, 0xfc000003, 0xe8000000, 0x0, // Load Doubleword DS-form (ld RT,DS(RA))
		[5]*argField{ap_Reg_6_10, ap_Offset_16_29_shift2, ap_Reg_11_15}},
	{LDU, 0xfc000003, 0xe8000001, 0x0, // Load Doubleword with Update DS-form (ldu RT,DS(RA))
		[5]*argField{ap_Reg_6_10, ap_Offset_16_29_shift2, ap_Reg_11_15}},
	{LDX, 0xfc0007fe, 0x7c00002a, 0x1, // Load Doubleword Indexed X-form (ldx RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LDUX, 0xfc0007fe, 0x7c00006a, 0x1, // Load Doubleword with Update Indexed X-form (ldux RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STB, 0xfc000000, 0x98000000, 0x0, // Store Byte D-form (stb RS,D(RA))
		[5]*argField{ap_Reg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{STBU, 0xfc000000, 0x9c000000, 0x0, // Store Byte with Update D-form (stbu RS,D(RA))
		[5]*argField{ap_Reg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{STBX, 0xfc0007fe, 0x7c0001ae, 0x1, // Store Byte Indexed X-form (stbx RS,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STBUX, 0xfc0007fe, 0x7c0001ee, 0x1, // Store Byte with Update Indexed X-form (stbux RS,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STH, 0xfc000000, 0xb0000000, 0x0, // Store Halfword D-form (sth RS,D(RA))
		[5]*argField{ap_Reg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{STHU, 0xfc000000, 0xb4000000, 0x0, // Store Halfword with Update D-form (sthu RS,D(RA))
		[5]*argField{ap_Reg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{STHX, 0xfc0007fe, 0x7c00032e, 0x1, // Store Halfword Indexed X-form (sthx RS,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STHUX, 0xfc0007fe, 0x7c00036e, 0x1, // Store Halfword with Update Indexed X-form (sthux RS,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STW, 0xfc000000, 0x90000000, 0x0, // Store Word D-form (stw RS,D(RA))
		[5]*argField{ap_Reg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{STWU, 0xfc000000, 0x94000000, 0x0, // Store Word with Update D-form (stwu RS,D(RA))
		[5]*argField{ap_Reg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{STWX, 0xfc0007fe, 0x7c00012e, 0x1, // Store Word Indexed X-form (stwx RS,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STWUX, 0xfc0007fe, 0x7c00016e, 0x1, // Store Word with Update Indexed X-form (stwux RS,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STD, 0xfc000003, 0xf8000000, 0x0, // Store Doubleword DS-form (std RS,DS(RA))
		[5]*argField{ap_Reg_6_10, ap_Offset_16_29_shift2, ap_Reg_11_15}},
	{STDU, 0xfc000003, 0xf8000001, 0x0, // Store Doubleword with Update DS-form (stdu RS,DS(RA))
		[5]*argField{ap_Reg_6_10, ap_Offset_16_29_shift2, ap_Reg_11_15}},
	{STDX, 0xfc0007fe, 0x7c00012a, 0x1, // Store Doubleword Indexed X-form (stdx RS,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STDUX, 0xfc0007fe, 0x7c00016a, 0x1, // Store Doubleword with Update Indexed X-form (stdux RS,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LQ, 0xfc000000, 0xe0000000, 0xf, // Load Quadword DQ-form (lq RTp,DQ(RA))
		[5]*argField{ap_Reg_6_10, ap_Offset_16_27_shift4, ap_Reg_11_15}},
	{STQ, 0xfc000003, 0xf8000002, 0x0, // Store Quadword DS-form (stq RSp,DS(RA))
		[5]*argField{ap_Reg_6_10, ap_Offset_16_29_shift2, ap_Reg_11_15}},
	{LHBRX, 0xfc0007fe, 0x7c00062c, 0x1, // Load Halfword Byte-Reverse Indexed X-form (lhbrx RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LWBRX, 0xfc0007fe, 0x7c00042c, 0x1, // Load Word Byte-Reverse Indexed X-form (lwbrx RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STHBRX, 0xfc0007fe, 0x7c00072c, 0x1, // Store Halfword Byte-Reverse Indexed X-form (sthbrx RS,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STWBRX, 0xfc0007fe, 0x7c00052c, 0x1, // Store Word Byte-Reverse Indexed X-form (stwbrx RS,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LDBRX, 0xfc0007fe, 0x7c000428, 0x1, // Load Doubleword Byte-Reverse Indexed X-form (ldbrx RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STDBRX, 0xfc0007fe, 0x7c000528, 0x1, // Store Doubleword Byte-Reverse Indexed X-form (stdbrx RS,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LMW, 0xfc000000, 0xb8000000, 0x0, // Load Multiple Word D-form (lmw RT,D(RA))
		[5]*argField{ap_Reg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{STMW, 0xfc000000, 0xbc000000, 0x0, // Store Multiple Word D-form (stmw RS,D(RA))
		[5]*argField{ap_Reg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{LSWI, 0xfc0007fe, 0x7c0004aa, 0x1, // Load String Word Immediate X-form (lswi RT,RA,NB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_ImmUnsigned_16_20}},
	{LSWX, 0xfc0007fe, 0x7c00042a, 0x1, // Load String Word Indexed X-form (lswx RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STSWI, 0xfc0007fe, 0x7c0005aa, 0x1, // Store String Word Immediate X-form (stswi RS,RA,NB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_ImmUnsigned_16_20}},
	{STSWX, 0xfc0007fe, 0x7c00052a, 0x1, // Store String Word Indexed X-form (stswx RS,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LI, 0xfc1f0000, 0x38000000, 0x0, // Add Immediate D-form (li RT,SI)
		[5]*argField{ap_Reg_6_10, ap_ImmSigned_16_31}},
	{ADDI, 0xfc000000, 0x38000000, 0x0, // Add Immediate D-form (addi RT,RA,SI)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_ImmSigned_16_31}},
	{LIS, 0xfc1f0000, 0x3c000000, 0x0, // Add Immediate Shifted D-form (lis RT, SI)
		[5]*argField{ap_Reg_6_10, ap_ImmSigned_16_31}},
	{ADDIS, 0xfc000000, 0x3c000000, 0x0, // Add Immediate Shifted D-form (addis RT,RA,SI)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_ImmSigned_16_31}},
	{ADD, 0xfc0007ff, 0x7c000214, 0x0, // Add XO-form (add RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ADD_, 0xfc0007ff, 0x7c000215, 0x0, // Add XO-form (add. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ADDO, 0xfc0007ff, 0x7c000614, 0x0, // Add XO-form (addo RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ADDO_, 0xfc0007ff, 0x7c000615, 0x0, // Add XO-form (addo. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ADDIC, 0xfc000000, 0x30000000, 0x0, // Add Immediate Carrying D-form (addic RT,RA,SI)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_ImmSigned_16_31}},
	{SUBF, 0xfc0007ff, 0x7c000050, 0x0, // Subtract From XO-form (subf RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{SUBF_, 0xfc0007ff, 0x7c000051, 0x0, // Subtract From XO-form (subf. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{SUBFO, 0xfc0007ff, 0x7c000450, 0x0, // Subtract From XO-form (subfo RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{SUBFO_, 0xfc0007ff, 0x7c000451, 0x0, // Subtract From XO-form (subfo. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ADDIC_, 0xfc000000, 0x34000000, 0x0, // Add Immediate Carrying and Record D-form (addic. RT,RA,SI)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_ImmSigned_16_31}},
	{SUBFIC, 0xfc000000, 0x20000000, 0x0, // Subtract From Immediate Carrying D-form (subfic RT,RA,SI)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_ImmSigned_16_31}},
	{ADDC, 0xfc0007ff, 0x7c000014, 0x0, // Add Carrying XO-form (addc RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ADDC_, 0xfc0007ff, 0x7c000015, 0x0, // Add Carrying XO-form (addc. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ADDCO, 0xfc0007ff, 0x7c000414, 0x0, // Add Carrying XO-form (addco RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ADDCO_, 0xfc0007ff, 0x7c000415, 0x0, // Add Carrying XO-form (addco. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{SUBFC, 0xfc0007ff, 0x7c000010, 0x0, // Subtract From Carrying XO-form (subfc RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{SUBFC_, 0xfc0007ff, 0x7c000011, 0x0, // Subtract From Carrying XO-form (subfc. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{SUBFCO, 0xfc0007ff, 0x7c000410, 0x0, // Subtract From Carrying XO-form (subfco RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{SUBFCO_, 0xfc0007ff, 0x7c000411, 0x0, // Subtract From Carrying XO-form (subfco. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ADDE, 0xfc0007ff, 0x7c000114, 0x0, // Add Extended XO-form (adde RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ADDE_, 0xfc0007ff, 0x7c000115, 0x0, // Add Extended XO-form (adde. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ADDEO, 0xfc0007ff, 0x7c000514, 0x0, // Add Extended XO-form (addeo RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ADDEO_, 0xfc0007ff, 0x7c000515, 0x0, // Add Extended XO-form (addeo. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ADDME, 0xfc0007ff, 0x7c0001d4, 0xf800, // Add to Minus One Extended XO-form (addme RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{ADDME_, 0xfc0007ff, 0x7c0001d5, 0xf800, // Add to Minus One Extended XO-form (addme. RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{ADDMEO, 0xfc0007ff, 0x7c0005d4, 0xf800, // Add to Minus One Extended XO-form (addmeo RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{ADDMEO_, 0xfc0007ff, 0x7c0005d5, 0xf800, // Add to Minus One Extended XO-form (addmeo. RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{SUBFE, 0xfc0007ff, 0x7c000110, 0x0, // Subtract From Extended XO-form (subfe RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{SUBFE_, 0xfc0007ff, 0x7c000111, 0x0, // Subtract From Extended XO-form (subfe. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{SUBFEO, 0xfc0007ff, 0x7c000510, 0x0, // Subtract From Extended XO-form (subfeo RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{SUBFEO_, 0xfc0007ff, 0x7c000511, 0x0, // Subtract From Extended XO-form (subfeo. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{SUBFME, 0xfc0007ff, 0x7c0001d0, 0xf800, // Subtract From Minus One Extended XO-form (subfme RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{SUBFME_, 0xfc0007ff, 0x7c0001d1, 0xf800, // Subtract From Minus One Extended XO-form (subfme. RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{SUBFMEO, 0xfc0007ff, 0x7c0005d0, 0xf800, // Subtract From Minus One Extended XO-form (subfmeo RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{SUBFMEO_, 0xfc0007ff, 0x7c0005d1, 0xf800, // Subtract From Minus One Extended XO-form (subfmeo. RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{ADDZE, 0xfc0007ff, 0x7c000194, 0xf800, // Add to Zero Extended XO-form (addze RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{ADDZE_, 0xfc0007ff, 0x7c000195, 0xf800, // Add to Zero Extended XO-form (addze. RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{ADDZEO, 0xfc0007ff, 0x7c000594, 0xf800, // Add to Zero Extended XO-form (addzeo RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{ADDZEO_, 0xfc0007ff, 0x7c000595, 0xf800, // Add to Zero Extended XO-form (addzeo. RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{SUBFZE, 0xfc0007ff, 0x7c000190, 0xf800, // Subtract From Zero Extended XO-form (subfze RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{SUBFZE_, 0xfc0007ff, 0x7c000191, 0xf800, // Subtract From Zero Extended XO-form (subfze. RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{SUBFZEO, 0xfc0007ff, 0x7c000590, 0xf800, // Subtract From Zero Extended XO-form (subfzeo RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{SUBFZEO_, 0xfc0007ff, 0x7c000591, 0xf800, // Subtract From Zero Extended XO-form (subfzeo. RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{NEG, 0xfc0007ff, 0x7c0000d0, 0xf800, // Negate XO-form (neg RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{NEG_, 0xfc0007ff, 0x7c0000d1, 0xf800, // Negate XO-form (neg. RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{NEGO, 0xfc0007ff, 0x7c0004d0, 0xf800, // Negate XO-form (nego RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{NEGO_, 0xfc0007ff, 0x7c0004d1, 0xf800, // Negate XO-form (nego. RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{MULLI, 0xfc000000, 0x1c000000, 0x0, // Multiply Low Immediate D-form (mulli RT,RA,SI)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_ImmSigned_16_31}},
	{MULLW, 0xfc0007ff, 0x7c0001d6, 0x0, // Multiply Low Word XO-form (mullw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULLW_, 0xfc0007ff, 0x7c0001d7, 0x0, // Multiply Low Word XO-form (mullw. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULLWO, 0xfc0007ff, 0x7c0005d6, 0x0, // Multiply Low Word XO-form (mullwo RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULLWO_, 0xfc0007ff, 0x7c0005d7, 0x0, // Multiply Low Word XO-form (mullwo. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULHW, 0xfc0003ff, 0x7c000096, 0x400, // Multiply High Word XO-form (mulhw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULHW_, 0xfc0003ff, 0x7c000097, 0x400, // Multiply High Word XO-form (mulhw. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULHWU, 0xfc0003ff, 0x7c000016, 0x400, // Multiply High Word Unsigned XO-form (mulhwu RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULHWU_, 0xfc0003ff, 0x7c000017, 0x400, // Multiply High Word Unsigned XO-form (mulhwu. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVW, 0xfc0007ff, 0x7c0003d6, 0x0, // Divide Word XO-form (divw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVW_, 0xfc0007ff, 0x7c0003d7, 0x0, // Divide Word XO-form (divw. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVWO, 0xfc0007ff, 0x7c0007d6, 0x0, // Divide Word XO-form (divwo RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVWO_, 0xfc0007ff, 0x7c0007d7, 0x0, // Divide Word XO-form (divwo. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVWU, 0xfc0007ff, 0x7c000396, 0x0, // Divide Word Unsigned XO-form (divwu RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVWU_, 0xfc0007ff, 0x7c000397, 0x0, // Divide Word Unsigned XO-form (divwu. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVWUO, 0xfc0007ff, 0x7c000796, 0x0, // Divide Word Unsigned XO-form (divwuo RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVWUO_, 0xfc0007ff, 0x7c000797, 0x0, // Divide Word Unsigned XO-form (divwuo. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVWE, 0xfc0007ff, 0x7c000356, 0x0, // Divide Word Extended XO-form (divwe RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVWE_, 0xfc0007ff, 0x7c000357, 0x0, // Divide Word Extended XO-form (divwe. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVWEO, 0xfc0007ff, 0x7c000756, 0x0, // Divide Word Extended XO-form (divweo RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVWEO_, 0xfc0007ff, 0x7c000757, 0x0, // Divide Word Extended XO-form (divweo. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVWEU, 0xfc0007ff, 0x7c000316, 0x0, // Divide Word Extended Unsigned XO-form (divweu RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVWEU_, 0xfc0007ff, 0x7c000317, 0x0, // Divide Word Extended Unsigned XO-form (divweu. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVWEUO, 0xfc0007ff, 0x7c000716, 0x0, // Divide Word Extended Unsigned XO-form (divweuo RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVWEUO_, 0xfc0007ff, 0x7c000717, 0x0, // Divide Word Extended Unsigned XO-form (divweuo. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULLD, 0xfc0007ff, 0x7c0001d2, 0x0, // Multiply Low Doubleword XO-form (mulld RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULLD_, 0xfc0007ff, 0x7c0001d3, 0x0, // Multiply Low Doubleword XO-form (mulld. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULLDO, 0xfc0007ff, 0x7c0005d2, 0x0, // Multiply Low Doubleword XO-form (mulldo RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULLDO_, 0xfc0007ff, 0x7c0005d3, 0x0, // Multiply Low Doubleword XO-form (mulldo. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULHDU, 0xfc0003ff, 0x7c000012, 0x400, // Multiply High Doubleword Unsigned XO-form (mulhdu RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULHDU_, 0xfc0003ff, 0x7c000013, 0x400, // Multiply High Doubleword Unsigned XO-form (mulhdu. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULHD, 0xfc0003ff, 0x7c000092, 0x400, // Multiply High Doubleword XO-form (mulhd RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULHD_, 0xfc0003ff, 0x7c000093, 0x400, // Multiply High Doubleword XO-form (mulhd. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVD, 0xfc0007ff, 0x7c0003d2, 0x0, // Divide Doubleword XO-form (divd RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVD_, 0xfc0007ff, 0x7c0003d3, 0x0, // Divide Doubleword XO-form (divd. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVDO, 0xfc0007ff, 0x7c0007d2, 0x0, // Divide Doubleword XO-form (divdo RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVDO_, 0xfc0007ff, 0x7c0007d3, 0x0, // Divide Doubleword XO-form (divdo. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVDU, 0xfc0007ff, 0x7c000392, 0x0, // Divide Doubleword Unsigned XO-form (divdu RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVDU_, 0xfc0007ff, 0x7c000393, 0x0, // Divide Doubleword Unsigned XO-form (divdu. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVDUO, 0xfc0007ff, 0x7c000792, 0x0, // Divide Doubleword Unsigned XO-form (divduo RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVDUO_, 0xfc0007ff, 0x7c000793, 0x0, // Divide Doubleword Unsigned XO-form (divduo. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVDE, 0xfc0007ff, 0x7c000352, 0x0, // Divide Doubleword Extended XO-form (divde RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVDE_, 0xfc0007ff, 0x7c000353, 0x0, // Divide Doubleword Extended XO-form (divde. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVDEO, 0xfc0007ff, 0x7c000752, 0x0, // Divide Doubleword Extended XO-form (divdeo RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVDEO_, 0xfc0007ff, 0x7c000753, 0x0, // Divide Doubleword Extended XO-form (divdeo. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVDEU, 0xfc0007ff, 0x7c000312, 0x0, // Divide Doubleword Extended Unsigned XO-form (divdeu RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVDEU_, 0xfc0007ff, 0x7c000313, 0x0, // Divide Doubleword Extended Unsigned XO-form (divdeu. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVDEUO, 0xfc0007ff, 0x7c000712, 0x0, // Divide Doubleword Extended Unsigned XO-form (divdeuo RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVDEUO_, 0xfc0007ff, 0x7c000713, 0x0, // Divide Doubleword Extended Unsigned XO-form (divdeuo. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{CMPWI, 0xfc200000, 0x2c000000, 0x400000, // Compare Immediate D-form (cmpwi BF,RA,SI)
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_ImmSigned_16_31}},
	{CMPDI, 0xfc200000, 0x2c200000, 0x400000, // Compare Immediate D-form (cmpdi BF,RA,SI)
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_ImmSigned_16_31}},
	{CMPW, 0xfc2007fe, 0x7c000000, 0x400001, // Compare X-form (cmpw BF,RA,RB)
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{CMPD, 0xfc2007fe, 0x7c200000, 0x400001, // Compare X-form (cmpd BF,RA,RB)
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{CMPLWI, 0xfc200000, 0x28000000, 0x400000, // Compare Logical Immediate D-form (cmplwi BF,RA,UI)
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_ImmUnsigned_16_31}},
	{CMPLDI, 0xfc200000, 0x28200000, 0x400000, // Compare Logical Immediate D-form (cmpldi BF,RA,UI)
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_ImmUnsigned_16_31}},
	{CMPLW, 0xfc2007fe, 0x7c000040, 0x400001, // Compare Logical X-form (cmplw BF,RA,RB)
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{CMPLD, 0xfc2007fe, 0x7c200040, 0x400001, // Compare Logical X-form (cmpld BF,RA,RB)
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{TWI, 0xfc000000, 0xc000000, 0x0, // Trap Word Immediate D-form (twi TO,RA,SI)
		[5]*argField{ap_ImmUnsigned_6_10, ap_Reg_11_15, ap_ImmSigned_16_31}},
	{TW, 0xfc0007fe, 0x7c000008, 0x1, // Trap Word X-form (tw TO,RA,RB)
		[5]*argField{ap_ImmUnsigned_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{TDI, 0xfc000000, 0x8000000, 0x0, // Trap Doubleword Immediate D-form (tdi TO,RA,SI)
		[5]*argField{ap_ImmUnsigned_6_10, ap_Reg_11_15, ap_ImmSigned_16_31}},
	{ISEL, 0xfc00003e, 0x7c00001e, 0x1, // Integer Select A-form (isel RT,RA,RB,BC)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20, ap_CondRegBit_21_25}},
	{TD, 0xfc0007fe, 0x7c000088, 0x1, // Trap Doubleword X-form (td TO,RA,RB)
		[5]*argField{ap_ImmUnsigned_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ANDI_, 0xfc000000, 0x70000000, 0x0, // AND Immediate D-form (andi. RA,RS,UI)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_16_31}},
	{ANDIS_, 0xfc000000, 0x74000000, 0x0, // AND Immediate Shifted D-form (andis. RA,RS,UI)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_16_31}},
	{ORI, 0xfc000000, 0x60000000, 0x0, // OR Immediate D-form (ori RA,RS,UI)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_16_31}},
	{ORIS, 0xfc000000, 0x64000000, 0x0, // OR Immediate Shifted D-form (oris RA,RS,UI)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_16_31}},
	{XORI, 0xfc000000, 0x68000000, 0x0, // XOR Immediate D-form (xori RA,RS,UI)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_16_31}},
	{XORIS, 0xfc000000, 0x6c000000, 0x0, // XOR Immediate Shifted D-form (xoris RA,RS,UI)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_16_31}},
	{AND, 0xfc0007ff, 0x7c000038, 0x0, // AND X-form (and RA,RS,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{AND_, 0xfc0007ff, 0x7c000039, 0x0, // AND X-form (and. RA,RS,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{XOR, 0xfc0007ff, 0x7c000278, 0x0, // XOR X-form (xor RA,RS,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{XOR_, 0xfc0007ff, 0x7c000279, 0x0, // XOR X-form (xor. RA,RS,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{NAND, 0xfc0007ff, 0x7c0003b8, 0x0, // NAND X-form (nand RA,RS,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{NAND_, 0xfc0007ff, 0x7c0003b9, 0x0, // NAND X-form (nand. RA,RS,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{OR, 0xfc0007ff, 0x7c000378, 0x0, // OR X-form (or RA,RS,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{OR_, 0xfc0007ff, 0x7c000379, 0x0, // OR X-form (or. RA,RS,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{NOR, 0xfc0007ff, 0x7c0000f8, 0x0, // NOR X-form (nor RA,RS,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{NOR_, 0xfc0007ff, 0x7c0000f9, 0x0, // NOR X-form (nor. RA,RS,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{ANDC, 0xfc0007ff, 0x7c000078, 0x0, // AND with Complement X-form (andc RA,RS,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{ANDC_, 0xfc0007ff, 0x7c000079, 0x0, // AND with Complement X-form (andc. RA,RS,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{EXTSB, 0xfc0007ff, 0x7c000774, 0xf800, // Extend Sign Byte X-form (extsb RA,RS)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10}},
	{EXTSB_, 0xfc0007ff, 0x7c000775, 0xf800, // Extend Sign Byte X-form (extsb. RA,RS)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10}},
	{EQV, 0xfc0007ff, 0x7c000238, 0x0, // Equivalent X-form (eqv RA,RS,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{EQV_, 0xfc0007ff, 0x7c000239, 0x0, // Equivalent X-form (eqv. RA,RS,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{ORC, 0xfc0007ff, 0x7c000338, 0x0, // OR with Complement X-form (orc RA,RS,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{ORC_, 0xfc0007ff, 0x7c000339, 0x0, // OR with Complement X-form (orc. RA,RS,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{EXTSH, 0xfc0007ff, 0x7c000734, 0xf800, // Extend Sign Halfword X-form (extsh RA,RS)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10}},
	{EXTSH_, 0xfc0007ff, 0x7c000735, 0xf800, // Extend Sign Halfword X-form (extsh. RA,RS)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10}},
	{CMPB, 0xfc0007fe, 0x7c0003f8, 0x1, // Compare Bytes X-form (cmpb RA,RS,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{POPCNTB, 0xfc0007fe, 0x7c0000f4, 0xf801, // Population Count Bytes X-form (popcntb RA, RS)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10}},
	{POPCNTW, 0xfc0007fe, 0x7c0002f4, 0xf801, // Population Count Words X-form (popcntw RA, RS)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10}},
	{PRTYD, 0xfc0007fe, 0x7c000174, 0xf801, // Parity Doubleword X-form (prtyd RA,RS)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10}},
	{PRTYW, 0xfc0007fe, 0x7c000134, 0xf801, // Parity Word X-form (prtyw RA,RS)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10}},
	{EXTSW, 0xfc0007ff, 0x7c0007b4, 0xf800, // Extend Sign Word X-form (extsw RA,RS)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10}},
	{EXTSW_, 0xfc0007ff, 0x7c0007b5, 0xf800, // Extend Sign Word X-form (extsw. RA,RS)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10}},
	{CNTLZD, 0xfc0007ff, 0x7c000074, 0xf800, // Count Leading Zeros Doubleword X-form (cntlzd RA,RS)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10}},
	{CNTLZD_, 0xfc0007ff, 0x7c000075, 0xf800, // Count Leading Zeros Doubleword X-form (cntlzd. RA,RS)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10}},
	{POPCNTD, 0xfc0007fe, 0x7c0003f4, 0xf801, // Population Count Doubleword X-form (popcntd RA, RS)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10}},
	{BPERMD, 0xfc0007fe, 0x7c0001f8, 0x1, // Bit Permute Doubleword X-form (bpermd RA,RS,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{RLWINM, 0xfc000001, 0x54000000, 0x0, // Rotate Left Word Immediate then AND with Mask M-form (rlwinm RA,RS,SH,MB,ME)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_ImmUnsigned_21_25, ap_ImmUnsigned_26_30}},
	{RLWINM_, 0xfc000001, 0x54000001, 0x0, // Rotate Left Word Immediate then AND with Mask M-form (rlwinm. RA,RS,SH,MB,ME)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_ImmUnsigned_21_25, ap_ImmUnsigned_26_30}},
	{RLWNM, 0xfc000001, 0x5c000000, 0x0, // Rotate Left Word then AND with Mask M-form (rlwnm RA,RS,RB,MB,ME)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20, ap_ImmUnsigned_21_25, ap_ImmUnsigned_26_30}},
	{RLWNM_, 0xfc000001, 0x5c000001, 0x0, // Rotate Left Word then AND with Mask M-form (rlwnm. RA,RS,RB,MB,ME)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20, ap_ImmUnsigned_21_25, ap_ImmUnsigned_26_30}},
	{RLWIMI, 0xfc000001, 0x50000000, 0x0, // Rotate Left Word Immediate then Mask Insert M-form (rlwimi RA,RS,SH,MB,ME)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_ImmUnsigned_21_25, ap_ImmUnsigned_26_30}},
	{RLWIMI_, 0xfc000001, 0x50000001, 0x0, // Rotate Left Word Immediate then Mask Insert M-form (rlwimi. RA,RS,SH,MB,ME)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_ImmUnsigned_21_25, ap_ImmUnsigned_26_30}},
	{RLDICL, 0xfc00001d, 0x78000000, 0x0, // Rotate Left Doubleword Immediate then Clear Left MD-form (rldicl RA,RS,SH,MB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_30_30_16_20, ap_ImmUnsigned_26_26_21_25}},
	{RLDICL_, 0xfc00001d, 0x78000001, 0x0, // Rotate Left Doubleword Immediate then Clear Left MD-form (rldicl. RA,RS,SH,MB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_30_30_16_20, ap_ImmUnsigned_26_26_21_25}},
	{RLDICR, 0xfc00001d, 0x78000004, 0x0, // Rotate Left Doubleword Immediate then Clear Right MD-form (rldicr RA,RS,SH,ME)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_30_30_16_20, ap_ImmUnsigned_26_26_21_25}},
	{RLDICR_, 0xfc00001d, 0x78000005, 0x0, // Rotate Left Doubleword Immediate then Clear Right MD-form (rldicr. RA,RS,SH,ME)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_30_30_16_20, ap_ImmUnsigned_26_26_21_25}},
	{RLDIC, 0xfc00001d, 0x78000008, 0x0, // Rotate Left Doubleword Immediate then Clear MD-form (rldic RA,RS,SH,MB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_30_30_16_20, ap_ImmUnsigned_26_26_21_25}},
	{RLDIC_, 0xfc00001d, 0x78000009, 0x0, // Rotate Left Doubleword Immediate then Clear MD-form (rldic. RA,RS,SH,MB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_30_30_16_20, ap_ImmUnsigned_26_26_21_25}},
	{RLDCL, 0xfc00001f, 0x78000010, 0x0, // Rotate Left Doubleword then Clear Left MDS-form (rldcl RA,RS,RB,MB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20, ap_ImmUnsigned_26_26_21_25}},
	{RLDCL_, 0xfc00001f, 0x78000011, 0x0, // Rotate Left Doubleword then Clear Left MDS-form (rldcl. RA,RS,RB,MB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20, ap_ImmUnsigned_26_26_21_25}},
	{RLDCR, 0xfc00001f, 0x78000012, 0x0, // Rotate Left Doubleword then Clear Right MDS-form (rldcr RA,RS,RB,ME)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20, ap_ImmUnsigned_26_26_21_25}},
	{RLDCR_, 0xfc00001f, 0x78000013, 0x0, // Rotate Left Doubleword then Clear Right MDS-form (rldcr. RA,RS,RB,ME)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20, ap_ImmUnsigned_26_26_21_25}},
	{RLDIMI, 0xfc00001d, 0x7800000c, 0x0, // Rotate Left Doubleword Immediate then Mask Insert MD-form (rldimi RA,RS,SH,MB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_30_30_16_20, ap_ImmUnsigned_26_26_21_25}},
	{RLDIMI_, 0xfc00001d, 0x7800000d, 0x0, // Rotate Left Doubleword Immediate then Mask Insert MD-form (rldimi. RA,RS,SH,MB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_30_30_16_20, ap_ImmUnsigned_26_26_21_25}},
	{SLW, 0xfc0007ff, 0x7c000030, 0x0, // Shift Left Word X-form (slw RA,RS,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{SLW_, 0xfc0007ff, 0x7c000031, 0x0, // Shift Left Word X-form (slw. RA,RS,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{SRW, 0xfc0007ff, 0x7c000430, 0x0, // Shift Right Word X-form (srw RA,RS,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{SRW_, 0xfc0007ff, 0x7c000431, 0x0, // Shift Right Word X-form (srw. RA,RS,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{SRAWI, 0xfc0007ff, 0x7c000670, 0x0, // Shift Right Algebraic Word Immediate X-form (srawi RA,RS,SH)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_16_20}},
	{SRAWI_, 0xfc0007ff, 0x7c000671, 0x0, // Shift Right Algebraic Word Immediate X-form (srawi. RA,RS,SH)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_16_20}},
	{SRAW, 0xfc0007ff, 0x7c000630, 0x0, // Shift Right Algebraic Word X-form (sraw RA,RS,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{SRAW_, 0xfc0007ff, 0x7c000631, 0x0, // Shift Right Algebraic Word X-form (sraw. RA,RS,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{SLD, 0xfc0007ff, 0x7c000036, 0x0, // Shift Left Doubleword X-form (sld RA,RS,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{SLD_, 0xfc0007ff, 0x7c000037, 0x0, // Shift Left Doubleword X-form (sld. RA,RS,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{SRD, 0xfc0007ff, 0x7c000436, 0x0, // Shift Right Doubleword X-form (srd RA,RS,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{SRD_, 0xfc0007ff, 0x7c000437, 0x0, // Shift Right Doubleword X-form (srd. RA,RS,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{SRADI, 0xfc0007fd, 0x7c000674, 0x0, // Shift Right Algebraic Doubleword Immediate XS-form (sradi RA,RS,SH)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_30_30_16_20}},
	{SRADI_, 0xfc0007fd, 0x7c000675, 0x0, // Shift Right Algebraic Doubleword Immediate XS-form (sradi. RA,RS,SH)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_30_30_16_20}},
	{SRAD, 0xfc0007ff, 0x7c000634, 0x0, // Shift Right Algebraic Doubleword X-form (srad RA,RS,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{SRAD_, 0xfc0007ff, 0x7c000635, 0x0, // Shift Right Algebraic Doubleword X-form (srad. RA,RS,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{CDTBCD, 0xfc0007fe, 0x7c000234, 0xf801, // Convert Declets To Binary Coded Decimal X-form (cdtbcd RA, RS)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10}},
	{CBCDTD, 0xfc0007fe, 0x7c000274, 0xf801, // Convert Binary Coded Decimal To Declets X-form (cbcdtd RA, RS)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10}},
	{ADDG6S, 0xfc0003fe, 0x7c000094, 0x401, // Add and Generate Sixes XO-form (addg6s RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MTSPR, 0xfc0007fe, 0x7c0003a6, 0x1, // Move To Special Purpose Register XFX-form (mtspr SPR,RS)
		[5]*argField{ap_SpReg_16_20_11_15, ap_Reg_6_10}},
	{MFSPR, 0xfc0007fe, 0x7c0002a6, 0x1, // Move From Special Purpose Register XFX-form (mfspr RT,SPR)
		[5]*argField{ap_Reg_6_10, ap_SpReg_16_20_11_15}},
	{MTCRF, 0xfc1007fe, 0x7c000120, 0x801, // Move To Condition Register Fields XFX-form (mtcrf FXM,RS)
		[5]*argField{ap_ImmUnsigned_12_19, ap_Reg_6_10}},
	{MFCR, 0xfc1007fe, 0x7c000026, 0xff801, // Move From Condition Register XFX-form (mfcr RT)
		[5]*argField{ap_Reg_6_10}},
	{MTSLE, 0xfc0007fe, 0x7c000126, 0x3dff801, // Move To Split Little Endian X-form (mtsle L)
		[5]*argField{ap_ImmUnsigned_10_10}},
	{MFVSRD, 0xfc0007fe, 0x7c000066, 0xf800, // Move From VSR Doubleword XX1-form (mfvsrd RA,XS)
		[5]*argField{ap_Reg_11_15, ap_VecReg_31_31_6_10}},
	{MFVSRWZ, 0xfc0007fe, 0x7c0000e6, 0xf800, // Move From VSR Word and Zero XX1-form (mfvsrwz RA,XS)
		[5]*argField{ap_Reg_11_15, ap_VecReg_31_31_6_10}},
	{MTVSRD, 0xfc0007fe, 0x7c000166, 0xf800, // Move To VSR Doubleword XX1-form (mtvsrd XT,RA)
		[5]*argField{ap_VecReg_31_31_6_10, ap_Reg_11_15}},
	{MTVSRWA, 0xfc0007fe, 0x7c0001a6, 0xf800, // Move To VSR Word Algebraic XX1-form (mtvsrwa XT,RA)
		[5]*argField{ap_VecReg_31_31_6_10, ap_Reg_11_15}},
	{MTVSRWZ, 0xfc0007fe, 0x7c0001e6, 0xf800, // Move To VSR Word and Zero XX1-form (mtvsrwz XT,RA)
		[5]*argField{ap_VecReg_31_31_6_10, ap_Reg_11_15}},
	{MTOCRF, 0xfc1007fe, 0x7c100120, 0x801, // Move To One Condition Register Field XFX-form (mtocrf FXM,RS)
		[5]*argField{ap_ImmUnsigned_12_19, ap_Reg_6_10}},
	{MFOCRF, 0xfc1007fe, 0x7c100026, 0x801, // Move From One Condition Register Field XFX-form (mfocrf RT,FXM)
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_12_19}},
	{MCRXR, 0xfc0007fe, 0x7c000400, 0x7ff801, // Move to Condition Register from XER X-form (mcrxr BF)
		[5]*argField{ap_CondRegField_6_8}},
	{MTDCRUX, 0xfc0007fe, 0x7c000346, 0xf801, // Move To Device Control Register User-mode Indexed X-form (mtdcrux RS,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{MFDCRUX, 0xfc0007fe, 0x7c000246, 0xf801, // Move From Device Control Register User-mode Indexed X-form (mfdcrux RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{LFS, 0xfc000000, 0xc0000000, 0x0, // Load Floating-Point Single D-form (lfs FRT,D(RA))
		[5]*argField{ap_FPReg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{LFSU, 0xfc000000, 0xc4000000, 0x0, // Load Floating-Point Single with Update D-form (lfsu FRT,D(RA))
		[5]*argField{ap_FPReg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{LFSX, 0xfc0007fe, 0x7c00042e, 0x1, // Load Floating-Point Single Indexed X-form (lfsx FRT,RA,RB)
		[5]*argField{ap_FPReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LFSUX, 0xfc0007fe, 0x7c00046e, 0x1, // Load Floating-Point Single with Update Indexed X-form (lfsux FRT,RA,RB)
		[5]*argField{ap_FPReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LFD, 0xfc000000, 0xc8000000, 0x0, // Load Floating-Point Double D-form (lfd FRT,D(RA))
		[5]*argField{ap_FPReg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{LFDU, 0xfc000000, 0xcc000000, 0x0, // Load Floating-Point Double with Update D-form (lfdu FRT,D(RA))
		[5]*argField{ap_FPReg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{LFDX, 0xfc0007fe, 0x7c0004ae, 0x1, // Load Floating-Point Double Indexed X-form (lfdx FRT,RA,RB)
		[5]*argField{ap_FPReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LFDUX, 0xfc0007fe, 0x7c0004ee, 0x1, // Load Floating-Point Double with Update Indexed X-form (lfdux FRT,RA,RB)
		[5]*argField{ap_FPReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LFIWAX, 0xfc0007fe, 0x7c0006ae, 0x1, // Load Floating-Point as Integer Word Algebraic Indexed X-form (lfiwax FRT,RA,RB)
		[5]*argField{ap_FPReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LFIWZX, 0xfc0007fe, 0x7c0006ee, 0x1, // Load Floating-Point as Integer Word and Zero Indexed X-form (lfiwzx FRT,RA,RB)
		[5]*argField{ap_FPReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STFS, 0xfc000000, 0xd0000000, 0x0, // Store Floating-Point Single D-form (stfs FRS,D(RA))
		[5]*argField{ap_FPReg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{STFSU, 0xfc000000, 0xd4000000, 0x0, // Store Floating-Point Single with Update D-form (stfsu FRS,D(RA))
		[5]*argField{ap_FPReg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{STFSX, 0xfc0007fe, 0x7c00052e, 0x1, // Store Floating-Point Single Indexed X-form (stfsx FRS,RA,RB)
		[5]*argField{ap_FPReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STFSUX, 0xfc0007fe, 0x7c00056e, 0x1, // Store Floating-Point Single with Update Indexed X-form (stfsux FRS,RA,RB)
		[5]*argField{ap_FPReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STFD, 0xfc000000, 0xd8000000, 0x0, // Store Floating-Point Double D-form (stfd FRS,D(RA))
		[5]*argField{ap_FPReg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{STFDU, 0xfc000000, 0xdc000000, 0x0, // Store Floating-Point Double with Update D-form (stfdu FRS,D(RA))
		[5]*argField{ap_FPReg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{STFDX, 0xfc0007fe, 0x7c0005ae, 0x1, // Store Floating-Point Double Indexed X-form (stfdx FRS,RA,RB)
		[5]*argField{ap_FPReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STFDUX, 0xfc0007fe, 0x7c0005ee, 0x1, // Store Floating-Point Double with Update Indexed X-form (stfdux FRS,RA,RB)
		[5]*argField{ap_FPReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STFIWX, 0xfc0007fe, 0x7c0007ae, 0x1, // Store Floating-Point as Integer Word Indexed X-form (stfiwx FRS,RA,RB)
		[5]*argField{ap_FPReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LFDP, 0xfc000003, 0xe4000000, 0x0, // Load Floating-Point Double Pair DS-form (lfdp FRTp,DS(RA))
		[5]*argField{ap_FPReg_6_10, ap_Offset_16_29_shift2, ap_Reg_11_15}},
	{LFDPX, 0xfc0007fe, 0x7c00062e, 0x1, // Load Floating-Point Double Pair Indexed X-form (lfdpx FRTp,RA,RB)
		[5]*argField{ap_FPReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STFDP, 0xfc000003, 0xf4000000, 0x0, // Store Floating-Point Double Pair DS-form (stfdp FRSp,DS(RA))
		[5]*argField{ap_FPReg_6_10, ap_Offset_16_29_shift2, ap_Reg_11_15}},
	{STFDPX, 0xfc0007fe, 0x7c00072e, 0x1, // Store Floating-Point Double Pair Indexed X-form (stfdpx FRSp,RA,RB)
		[5]*argField{ap_FPReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{FMR, 0xfc0007ff, 0xfc000090, 0x1f0000, // Floating Move Register X-form (fmr FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FMR_, 0xfc0007ff, 0xfc000091, 0x1f0000, // Floating Move Register X-form (fmr. FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FABS, 0xfc0007ff, 0xfc000210, 0x1f0000, // Floating Absolute Value X-form (fabs FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FABS_, 0xfc0007ff, 0xfc000211, 0x1f0000, // Floating Absolute Value X-form (fabs. FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FNABS, 0xfc0007ff, 0xfc000110, 0x1f0000, // Floating Negative Absolute Value X-form (fnabs FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FNABS_, 0xfc0007ff, 0xfc000111, 0x1f0000, // Floating Negative Absolute Value X-form (fnabs. FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FNEG, 0xfc0007ff, 0xfc000050, 0x1f0000, // Floating Negate X-form (fneg FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FNEG_, 0xfc0007ff, 0xfc000051, 0x1f0000, // Floating Negate X-form (fneg. FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCPSGN, 0xfc0007ff, 0xfc000010, 0x0, // Floating Copy Sign X-form (fcpsgn FRT, FRA, FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FCPSGN_, 0xfc0007ff, 0xfc000011, 0x0, // Floating Copy Sign X-form (fcpsgn. FRT, FRA, FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FMRGEW, 0xfc0007fe, 0xfc00078c, 0x1, // Floating Merge Even Word X-form (fmrgew FRT,FRA,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FMRGOW, 0xfc0007fe, 0xfc00068c, 0x1, // Floating Merge Odd Word X-form (fmrgow FRT,FRA,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FADD, 0xfc00003f, 0xfc00002a, 0x7c0, // Floating Add [Single] A-form (fadd FRT,FRA,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FADD_, 0xfc00003f, 0xfc00002b, 0x7c0, // Floating Add [Single] A-form (fadd. FRT,FRA,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FADDS, 0xfc00003f, 0xec00002a, 0x7c0, // Floating Add [Single] A-form (fadds FRT,FRA,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FADDS_, 0xfc00003f, 0xec00002b, 0x7c0, // Floating Add [Single] A-form (fadds. FRT,FRA,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FSUB, 0xfc00003f, 0xfc000028, 0x7c0, // Floating Subtract [Single] A-form (fsub FRT,FRA,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FSUB_, 0xfc00003f, 0xfc000029, 0x7c0, // Floating Subtract [Single] A-form (fsub. FRT,FRA,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FSUBS, 0xfc00003f, 0xec000028, 0x7c0, // Floating Subtract [Single] A-form (fsubs FRT,FRA,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FSUBS_, 0xfc00003f, 0xec000029, 0x7c0, // Floating Subtract [Single] A-form (fsubs. FRT,FRA,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FMUL, 0xfc00003f, 0xfc000032, 0xf800, // Floating Multiply [Single] A-form (fmul FRT,FRA,FRC)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25}},
	{FMUL_, 0xfc00003f, 0xfc000033, 0xf800, // Floating Multiply [Single] A-form (fmul. FRT,FRA,FRC)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25}},
	{FMULS, 0xfc00003f, 0xec000032, 0xf800, // Floating Multiply [Single] A-form (fmuls FRT,FRA,FRC)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25}},
	{FMULS_, 0xfc00003f, 0xec000033, 0xf800, // Floating Multiply [Single] A-form (fmuls. FRT,FRA,FRC)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25}},
	{FDIV, 0xfc00003f, 0xfc000024, 0x7c0, // Floating Divide [Single] A-form (fdiv FRT,FRA,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FDIV_, 0xfc00003f, 0xfc000025, 0x7c0, // Floating Divide [Single] A-form (fdiv. FRT,FRA,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FDIVS, 0xfc00003f, 0xec000024, 0x7c0, // Floating Divide [Single] A-form (fdivs FRT,FRA,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FDIVS_, 0xfc00003f, 0xec000025, 0x7c0, // Floating Divide [Single] A-form (fdivs. FRT,FRA,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FSQRT, 0xfc00003f, 0xfc00002c, 0x1f07c0, // Floating Square Root [Single] A-form (fsqrt FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FSQRT_, 0xfc00003f, 0xfc00002d, 0x1f07c0, // Floating Square Root [Single] A-form (fsqrt. FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FSQRTS, 0xfc00003f, 0xec00002c, 0x1f07c0, // Floating Square Root [Single] A-form (fsqrts FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FSQRTS_, 0xfc00003f, 0xec00002d, 0x1f07c0, // Floating Square Root [Single] A-form (fsqrts. FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FRE, 0xfc00003f, 0xfc000030, 0x1f07c0, // Floating Reciprocal Estimate [Single] A-form (fre FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FRE_, 0xfc00003f, 0xfc000031, 0x1f07c0, // Floating Reciprocal Estimate [Single] A-form (fre. FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FRES, 0xfc00003f, 0xec000030, 0x1f07c0, // Floating Reciprocal Estimate [Single] A-form (fres FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FRES_, 0xfc00003f, 0xec000031, 0x1f07c0, // Floating Reciprocal Estimate [Single] A-form (fres. FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FRSQRTE, 0xfc00003f, 0xfc000034, 0x1f07c0, // Floating Reciprocal Square Root Estimate [Single] A-form (frsqrte FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FRSQRTE_, 0xfc00003f, 0xfc000035, 0x1f07c0, // Floating Reciprocal Square Root Estimate [Single] A-form (frsqrte. FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FRSQRTES, 0xfc00003f, 0xec000034, 0x1f07c0, // Floating Reciprocal Square Root Estimate [Single] A-form (frsqrtes FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FRSQRTES_, 0xfc00003f, 0xec000035, 0x1f07c0, // Floating Reciprocal Square Root Estimate [Single] A-form (frsqrtes. FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FTDIV, 0xfc0007fe, 0xfc000100, 0x600001, // Floating Test for software Divide X-form (ftdiv BF,FRA,FRB)
		[5]*argField{ap_CondRegField_6_8, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FTSQRT, 0xfc0007fe, 0xfc000140, 0x7f0001, // Floating Test for software Square Root X-form (ftsqrt BF,FRB)
		[5]*argField{ap_CondRegField_6_8, ap_FPReg_16_20}},
	{FMADD, 0xfc00003f, 0xfc00003a, 0x0, // Floating Multiply-Add [Single] A-form (fmadd FRT,FRA,FRC,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25, ap_FPReg_16_20}},
	{FMADD_, 0xfc00003f, 0xfc00003b, 0x0, // Floating Multiply-Add [Single] A-form (fmadd. FRT,FRA,FRC,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25, ap_FPReg_16_20}},
	{FMADDS, 0xfc00003f, 0xec00003a, 0x0, // Floating Multiply-Add [Single] A-form (fmadds FRT,FRA,FRC,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25, ap_FPReg_16_20}},
	{FMADDS_, 0xfc00003f, 0xec00003b, 0x0, // Floating Multiply-Add [Single] A-form (fmadds. FRT,FRA,FRC,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25, ap_FPReg_16_20}},
	{FMSUB, 0xfc00003f, 0xfc000038, 0x0, // Floating Multiply-Subtract [Single] A-form (fmsub FRT,FRA,FRC,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25, ap_FPReg_16_20}},
	{FMSUB_, 0xfc00003f, 0xfc000039, 0x0, // Floating Multiply-Subtract [Single] A-form (fmsub. FRT,FRA,FRC,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25, ap_FPReg_16_20}},
	{FMSUBS, 0xfc00003f, 0xec000038, 0x0, // Floating Multiply-Subtract [Single] A-form (fmsubs FRT,FRA,FRC,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25, ap_FPReg_16_20}},
	{FMSUBS_, 0xfc00003f, 0xec000039, 0x0, // Floating Multiply-Subtract [Single] A-form (fmsubs. FRT,FRA,FRC,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25, ap_FPReg_16_20}},
	{FNMADD, 0xfc00003f, 0xfc00003e, 0x0, // Floating Negative Multiply-Add [Single] A-form (fnmadd FRT,FRA,FRC,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25, ap_FPReg_16_20}},
	{FNMADD_, 0xfc00003f, 0xfc00003f, 0x0, // Floating Negative Multiply-Add [Single] A-form (fnmadd. FRT,FRA,FRC,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25, ap_FPReg_16_20}},
	{FNMADDS, 0xfc00003f, 0xec00003e, 0x0, // Floating Negative Multiply-Add [Single] A-form (fnmadds FRT,FRA,FRC,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25, ap_FPReg_16_20}},
	{FNMADDS_, 0xfc00003f, 0xec00003f, 0x0, // Floating Negative Multiply-Add [Single] A-form (fnmadds. FRT,FRA,FRC,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25, ap_FPReg_16_20}},
	{FNMSUB, 0xfc00003f, 0xfc00003c, 0x0, // Floating Negative Multiply-Subtract [Single] A-form (fnmsub FRT,FRA,FRC,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25, ap_FPReg_16_20}},
	{FNMSUB_, 0xfc00003f, 0xfc00003d, 0x0, // Floating Negative Multiply-Subtract [Single] A-form (fnmsub. FRT,FRA,FRC,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25, ap_FPReg_16_20}},
	{FNMSUBS, 0xfc00003f, 0xec00003c, 0x0, // Floating Negative Multiply-Subtract [Single] A-form (fnmsubs FRT,FRA,FRC,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25, ap_FPReg_16_20}},
	{FNMSUBS_, 0xfc00003f, 0xec00003d, 0x0, // Floating Negative Multiply-Subtract [Single] A-form (fnmsubs. FRT,FRA,FRC,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25, ap_FPReg_16_20}},
	{FRSP, 0xfc0007ff, 0xfc000018, 0x1f0000, // Floating Round to Single-Precision X-form (frsp FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FRSP_, 0xfc0007ff, 0xfc000019, 0x1f0000, // Floating Round to Single-Precision X-form (frsp. FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCTID, 0xfc0007ff, 0xfc00065c, 0x1f0000, // Floating Convert To Integer Doubleword X-form (fctid FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCTID_, 0xfc0007ff, 0xfc00065d, 0x1f0000, // Floating Convert To Integer Doubleword X-form (fctid. FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCTIDZ, 0xfc0007ff, 0xfc00065e, 0x1f0000, // Floating Convert To Integer Doubleword with round toward Zero X-form (fctidz FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCTIDZ_, 0xfc0007ff, 0xfc00065f, 0x1f0000, // Floating Convert To Integer Doubleword with round toward Zero X-form (fctidz. FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCTIDU, 0xfc0007ff, 0xfc00075c, 0x1f0000, // Floating Convert To Integer Doubleword Unsigned X-form (fctidu FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCTIDU_, 0xfc0007ff, 0xfc00075d, 0x1f0000, // Floating Convert To Integer Doubleword Unsigned X-form (fctidu. FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCTIDUZ, 0xfc0007ff, 0xfc00075e, 0x1f0000, // Floating Convert To Integer Doubleword Unsigned with round toward Zero X-form (fctiduz FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCTIDUZ_, 0xfc0007ff, 0xfc00075f, 0x1f0000, // Floating Convert To Integer Doubleword Unsigned with round toward Zero X-form (fctiduz. FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCTIW, 0xfc0007ff, 0xfc00001c, 0x1f0000, // Floating Convert To Integer Word X-form (fctiw FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCTIW_, 0xfc0007ff, 0xfc00001d, 0x1f0000, // Floating Convert To Integer Word X-form (fctiw. FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCTIWZ, 0xfc0007ff, 0xfc00001e, 0x1f0000, // Floating Convert To Integer Word with round toward Zero X-form (fctiwz FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCTIWZ_, 0xfc0007ff, 0xfc00001f, 0x1f0000, // Floating Convert To Integer Word with round toward Zero X-form (fctiwz. FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCTIWU, 0xfc0007ff, 0xfc00011c, 0x1f0000, // Floating Convert To Integer Word Unsigned X-form (fctiwu FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCTIWU_, 0xfc0007ff, 0xfc00011d, 0x1f0000, // Floating Convert To Integer Word Unsigned X-form (fctiwu. FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCTIWUZ, 0xfc0007ff, 0xfc00011e, 0x1f0000, // Floating Convert To Integer Word Unsigned with round toward Zero X-form (fctiwuz FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCTIWUZ_, 0xfc0007ff, 0xfc00011f, 0x1f0000, // Floating Convert To Integer Word Unsigned with round toward Zero X-form (fctiwuz. FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCFID, 0xfc0007ff, 0xfc00069c, 0x1f0000, // Floating Convert From Integer Doubleword X-form (fcfid FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCFID_, 0xfc0007ff, 0xfc00069d, 0x1f0000, // Floating Convert From Integer Doubleword X-form (fcfid. FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCFIDU, 0xfc0007ff, 0xfc00079c, 0x1f0000, // Floating Convert From Integer Doubleword Unsigned X-form (fcfidu FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCFIDU_, 0xfc0007ff, 0xfc00079d, 0x1f0000, // Floating Convert From Integer Doubleword Unsigned X-form (fcfidu. FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCFIDS, 0xfc0007ff, 0xec00069c, 0x1f0000, // Floating Convert From Integer Doubleword Single X-form (fcfids FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCFIDS_, 0xfc0007ff, 0xec00069d, 0x1f0000, // Floating Convert From Integer Doubleword Single X-form (fcfids. FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCFIDUS, 0xfc0007ff, 0xec00079c, 0x1f0000, // Floating Convert From Integer Doubleword Unsigned Single X-form (fcfidus FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCFIDUS_, 0xfc0007ff, 0xec00079d, 0x1f0000, // Floating Convert From Integer Doubleword Unsigned Single X-form (fcfidus. FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FRIN, 0xfc0007ff, 0xfc000310, 0x1f0000, // Floating Round to Integer Nearest X-form (frin FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FRIN_, 0xfc0007ff, 0xfc000311, 0x1f0000, // Floating Round to Integer Nearest X-form (frin. FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FRIZ, 0xfc0007ff, 0xfc000350, 0x1f0000, // Floating Round to Integer Toward Zero X-form (friz FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FRIZ_, 0xfc0007ff, 0xfc000351, 0x1f0000, // Floating Round to Integer Toward Zero X-form (friz. FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FRIP, 0xfc0007ff, 0xfc000390, 0x1f0000, // Floating Round to Integer Plus X-form (frip FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FRIP_, 0xfc0007ff, 0xfc000391, 0x1f0000, // Floating Round to Integer Plus X-form (frip. FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FRIM, 0xfc0007ff, 0xfc0003d0, 0x1f0000, // Floating Round to Integer Minus X-form (frim FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FRIM_, 0xfc0007ff, 0xfc0003d1, 0x1f0000, // Floating Round to Integer Minus X-form (frim. FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCMPU, 0xfc0007fe, 0xfc000000, 0x600001, // Floating Compare Unordered X-form (fcmpu BF,FRA,FRB)
		[5]*argField{ap_CondRegField_6_8, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FCMPO, 0xfc0007fe, 0xfc000040, 0x600001, // Floating Compare Ordered X-form (fcmpo BF,FRA,FRB)
		[5]*argField{ap_CondRegField_6_8, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FSEL, 0xfc00003f, 0xfc00002e, 0x0, // Floating Select A-form (fsel FRT,FRA,FRC,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25, ap_FPReg_16_20}},
	{FSEL_, 0xfc00003f, 0xfc00002f, 0x0, // Floating Select A-form (fsel. FRT,FRA,FRC,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25, ap_FPReg_16_20}},
	{MFFS, 0xfc0007ff, 0xfc00048e, 0x1ff800, // Move From FPSCR X-form (mffs FRT)
		[5]*argField{ap_FPReg_6_10}},
	{MFFS_, 0xfc0007ff, 0xfc00048f, 0x1ff800, // Move From FPSCR X-form (mffs. FRT)
		[5]*argField{ap_FPReg_6_10}},
	{MCRFS, 0xfc0007fe, 0xfc000080, 0x63f801, // Move to Condition Register from FPSCR X-form (mcrfs BF,BFA)
		[5]*argField{ap_CondRegField_6_8, ap_CondRegField_11_13}},
	{MTFSFI, 0xfc0007ff, 0xfc00010c, 0x7e0800, // Move To FPSCR Field Immediate X-form (mtfsfi BF,U,W)
		[5]*argField{ap_CondRegField_6_8, ap_ImmUnsigned_16_19, ap_ImmUnsigned_15_15}},
	{MTFSFI_, 0xfc0007ff, 0xfc00010d, 0x7e0800, // Move To FPSCR Field Immediate X-form (mtfsfi. BF,U,W)
		[5]*argField{ap_CondRegField_6_8, ap_ImmUnsigned_16_19, ap_ImmUnsigned_15_15}},
	{MTFSF, 0xfc0007ff, 0xfc00058e, 0x0, // Move To FPSCR Fields XFL-form (mtfsf FLM,FRB,L,W)
		[5]*argField{ap_ImmUnsigned_7_14, ap_FPReg_16_20, ap_ImmUnsigned_6_6, ap_ImmUnsigned_15_15}},
	{MTFSF_, 0xfc0007ff, 0xfc00058f, 0x0, // Move To FPSCR Fields XFL-form (mtfsf. FLM,FRB,L,W)
		[5]*argField{ap_ImmUnsigned_7_14, ap_FPReg_16_20, ap_ImmUnsigned_6_6, ap_ImmUnsigned_15_15}},
	{MTFSB0, 0xfc0007ff, 0xfc00008c, 0x1ff800, // Move To FPSCR Bit 0 X-form (mtfsb0 BT)
		[5]*argField{ap_CondRegBit_6_10}},
	{MTFSB0_, 0xfc0007ff, 0xfc00008d, 0x1ff800, // Move To FPSCR Bit 0 X-form (mtfsb0. BT)
		[5]*argField{ap_CondRegBit_6_10}},
	{MTFSB1, 0xfc0007ff, 0xfc00004c, 0x1ff800, // Move To FPSCR Bit 1 X-form (mtfsb1 BT)
		[5]*argField{ap_CondRegBit_6_10}},
	{MTFSB1_, 0xfc0007ff, 0xfc00004d, 0x1ff800, // Move To FPSCR Bit 1 X-form (mtfsb1. BT)
		[5]*argField{ap_CondRegBit_6_10}},
	{LVEBX, 0xfc0007fe, 0x7c00000e, 0x1, // Load Vector Element Byte Indexed X-form (lvebx VRT,RA,RB)
		[5]*argField{ap_VecReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LVEHX, 0xfc0007fe, 0x7c00004e, 0x1, // Load Vector Element Halfword Indexed X-form (lvehx VRT,RA,RB)
		[5]*argField{ap_VecReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LVEWX, 0xfc0007fe, 0x7c00008e, 0x1, // Load Vector Element Word Indexed X-form (lvewx VRT,RA,RB)
		[5]*argField{ap_VecReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LVX, 0xfc0007fe, 0x7c0000ce, 0x1, // Load Vector Indexed X-form (lvx VRT,RA,RB)
		[5]*argField{ap_VecReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LVXL, 0xfc0007fe, 0x7c0002ce, 0x1, // Load Vector Indexed LRU X-form (lvxl VRT,RA,RB)
		[5]*argField{ap_VecReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STVEBX, 0xfc0007fe, 0x7c00010e, 0x1, // Store Vector Element Byte Indexed X-form (stvebx VRS,RA,RB)
		[5]*argField{ap_VecReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STVEHX, 0xfc0007fe, 0x7c00014e, 0x1, // Store Vector Element Halfword Indexed X-form (stvehx VRS,RA,RB)
		[5]*argField{ap_VecReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STVEWX, 0xfc0007fe, 0x7c00018e, 0x1, // Store Vector Element Word Indexed X-form (stvewx VRS,RA,RB)
		[5]*argField{ap_VecReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STVX, 0xfc0007fe, 0x7c0001ce, 0x1, // Store Vector Indexed X-form (stvx VRS,RA,RB)
		[5]*argField{ap_VecReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STVXL, 0xfc0007fe, 0x7c0003ce, 0x1, // Store Vector Indexed LRU X-form (stvxl VRS,RA,RB)
		[5]*argField{ap_VecReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LVSL, 0xfc0007fe, 0x7c00000c, 0x1, // Load Vector for Shift Left Indexed X-form (lvsl VRT,RA,RB)
		[5]*argField{ap_VecReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LVSR, 0xfc0007fe, 0x7c00004c, 0x1, // Load Vector for Shift Right Indexed X-form (lvsr VRT,RA,RB)
		[5]*argField{ap_VecReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{VPKPX, 0xfc0007ff, 0x1000030e, 0x0, // Vector Pack Pixel VX-form (vpkpx VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VPKSDSS, 0xfc0007ff, 0x100005ce, 0x0, // Vector Pack Signed Doubleword Signed Saturate VX-form (vpksdss VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VPKSDUS, 0xfc0007ff, 0x1000054e, 0x0, // Vector Pack Signed Doubleword Unsigned Saturate VX-form (vpksdus VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VPKSHSS, 0xfc0007ff, 0x1000018e, 0x0, // Vector Pack Signed Halfword Signed Saturate VX-form (vpkshss VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VPKSHUS, 0xfc0007ff, 0x1000010e, 0x0, // Vector Pack Signed Halfword Unsigned Saturate VX-form (vpkshus VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VPKSWSS, 0xfc0007ff, 0x100001ce, 0x0, // Vector Pack Signed Word Signed Saturate VX-form (vpkswss VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VPKSWUS, 0xfc0007ff, 0x1000014e, 0x0, // Vector Pack Signed Word Unsigned Saturate VX-form (vpkswus VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VPKUDUM, 0xfc0007ff, 0x1000044e, 0x0, // Vector Pack Unsigned Doubleword Unsigned Modulo VX-form (vpkudum VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VPKUDUS, 0xfc0007ff, 0x100004ce, 0x0, // Vector Pack Unsigned Doubleword Unsigned Saturate VX-form (vpkudus VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VPKUHUM, 0xfc0007ff, 0x1000000e, 0x0, // Vector Pack Unsigned Halfword Unsigned Modulo VX-form (vpkuhum VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VPKUHUS, 0xfc0007ff, 0x1000008e, 0x0, // Vector Pack Unsigned Halfword Unsigned Saturate VX-form (vpkuhus VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VPKUWUM, 0xfc0007ff, 0x1000004e, 0x0, // Vector Pack Unsigned Word Unsigned Modulo VX-form (vpkuwum VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VPKUWUS, 0xfc0007ff, 0x100000ce, 0x0, // Vector Pack Unsigned Word Unsigned Saturate VX-form (vpkuwus VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VUPKHPX, 0xfc0007ff, 0x1000034e, 0x1f0000, // Vector Unpack High Pixel VX-form (vupkhpx VRT,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VUPKLPX, 0xfc0007ff, 0x100003ce, 0x1f0000, // Vector Unpack Low Pixel VX-form (vupklpx VRT,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VUPKHSB, 0xfc0007ff, 0x1000020e, 0x1f0000, // Vector Unpack High Signed Byte VX-form (vupkhsb VRT,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VUPKHSH, 0xfc0007ff, 0x1000024e, 0x1f0000, // Vector Unpack High Signed Halfword VX-form (vupkhsh VRT,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VUPKHSW, 0xfc0007ff, 0x1000064e, 0x1f0000, // Vector Unpack High Signed Word VX-form (vupkhsw VRT,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VUPKLSB, 0xfc0007ff, 0x1000028e, 0x1f0000, // Vector Unpack Low Signed Byte VX-form (vupklsb VRT,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VUPKLSH, 0xfc0007ff, 0x100002ce, 0x1f0000, // Vector Unpack Low Signed Halfword VX-form (vupklsh VRT,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VUPKLSW, 0xfc0007ff, 0x100006ce, 0x1f0000, // Vector Unpack Low Signed Word VX-form (vupklsw VRT,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VMRGHB, 0xfc0007ff, 0x1000000c, 0x0, // Vector Merge High Byte VX-form (vmrghb VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMRGHH, 0xfc0007ff, 0x1000004c, 0x0, // Vector Merge High Halfword VX-form (vmrghh VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMRGLB, 0xfc0007ff, 0x1000010c, 0x0, // Vector Merge Low Byte VX-form (vmrglb VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMRGLH, 0xfc0007ff, 0x1000014c, 0x0, // Vector Merge Low Halfword VX-form (vmrglh VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMRGHW, 0xfc0007ff, 0x1000008c, 0x0, // Vector Merge High Word VX-form (vmrghw VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMRGLW, 0xfc0007ff, 0x1000018c, 0x0, // Vector Merge Low Word VX-form (vmrglw VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMRGEW, 0xfc0007ff, 0x1000078c, 0x0, // Vector Merge Even Word VX-form (vmrgew VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMRGOW, 0xfc0007ff, 0x1000068c, 0x0, // Vector Merge Odd Word VX-form (vmrgow VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSPLTB, 0xfc0007ff, 0x1000020c, 0x100000, // Vector Splat Byte VX-form (vspltb VRT,VRB,UIM)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20, ap_ImmUnsigned_12_15}},
	{VSPLTH, 0xfc0007ff, 0x1000024c, 0x180000, // Vector Splat Halfword VX-form (vsplth VRT,VRB,UIM)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20, ap_ImmUnsigned_13_15}},
	{VSPLTW, 0xfc0007ff, 0x1000028c, 0x1c0000, // Vector Splat Word VX-form (vspltw VRT,VRB,UIM)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20, ap_ImmUnsigned_14_15}},
	{VSPLTISB, 0xfc0007ff, 0x1000030c, 0xf800, // Vector Splat Immediate Signed Byte VX-form (vspltisb VRT,SIM)
		[5]*argField{ap_VecReg_6_10, ap_ImmSigned_11_15}},
	{VSPLTISH, 0xfc0007ff, 0x1000034c, 0xf800, // Vector Splat Immediate Signed Halfword VX-form (vspltish VRT,SIM)
		[5]*argField{ap_VecReg_6_10, ap_ImmSigned_11_15}},
	{VSPLTISW, 0xfc0007ff, 0x1000038c, 0xf800, // Vector Splat Immediate Signed Word VX-form (vspltisw VRT,SIM)
		[5]*argField{ap_VecReg_6_10, ap_ImmSigned_11_15}},
	{VPERM, 0xfc00003f, 0x1000002b, 0x0, // Vector Permute VA-form (vperm VRT,VRA,VRB,VRC)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_VecReg_21_25}},
	{VSEL, 0xfc00003f, 0x1000002a, 0x0, // Vector Select VA-form (vsel VRT,VRA,VRB,VRC)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_VecReg_21_25}},
	{VSL, 0xfc0007ff, 0x100001c4, 0x0, // Vector Shift Left VX-form (vsl VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSLDOI, 0xfc00003f, 0x1000002c, 0x400, // Vector Shift Left Double by Octet Immediate VA-form (vsldoi VRT,VRA,VRB,SHB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_ImmUnsigned_22_25}},
	{VSLO, 0xfc0007ff, 0x1000040c, 0x0, // Vector Shift Left by Octet VX-form (vslo VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSR, 0xfc0007ff, 0x100002c4, 0x0, // Vector Shift Right VX-form (vsr VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSRO, 0xfc0007ff, 0x1000044c, 0x0, // Vector Shift Right by Octet VX-form (vsro VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VADDCUW, 0xfc0007ff, 0x10000180, 0x0, // Vector Add and Write Carry-Out Unsigned Word VX-form (vaddcuw VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VADDSBS, 0xfc0007ff, 0x10000300, 0x0, // Vector Add Signed Byte Saturate VX-form (vaddsbs VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VADDSHS, 0xfc0007ff, 0x10000340, 0x0, // Vector Add Signed Halfword Saturate VX-form (vaddshs VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VADDSWS, 0xfc0007ff, 0x10000380, 0x0, // Vector Add Signed Word Saturate VX-form (vaddsws VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VADDUBM, 0xfc0007ff, 0x10000000, 0x0, // Vector Add Unsigned Byte Modulo VX-form (vaddubm VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VADDUDM, 0xfc0007ff, 0x100000c0, 0x0, // Vector Add Unsigned Doubleword Modulo VX-form (vaddudm VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VADDUHM, 0xfc0007ff, 0x10000040, 0x0, // Vector Add Unsigned Halfword Modulo VX-form (vadduhm VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VADDUWM, 0xfc0007ff, 0x10000080, 0x0, // Vector Add Unsigned Word Modulo VX-form (vadduwm VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VADDUBS, 0xfc0007ff, 0x10000200, 0x0, // Vector Add Unsigned Byte Saturate VX-form (vaddubs VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VADDUHS, 0xfc0007ff, 0x10000240, 0x0, // Vector Add Unsigned Halfword Saturate VX-form (vadduhs VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VADDUWS, 0xfc0007ff, 0x10000280, 0x0, // Vector Add Unsigned Word Saturate VX-form (vadduws VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VADDUQM, 0xfc0007ff, 0x10000100, 0x0, // Vector Add Unsigned Quadword Modulo VX-form (vadduqm VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VADDEUQM, 0xfc00003f, 0x1000003c, 0x0, // Vector Add Extended Unsigned Quadword Modulo VA-form (vaddeuqm VRT,VRA,VRB,VRC)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_VecReg_21_25}},
	{VADDCUQ, 0xfc0007ff, 0x10000140, 0x0, // Vector Add & write Carry Unsigned Quadword VX-form (vaddcuq VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VADDECUQ, 0xfc00003f, 0x1000003d, 0x0, // Vector Add Extended & write Carry Unsigned Quadword VA-form (vaddecuq VRT,VRA,VRB,VRC)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_VecReg_21_25}},
	{VSUBCUW, 0xfc0007ff, 0x10000580, 0x0, // Vector Subtract and Write Carry-Out Unsigned Word VX-form (vsubcuw VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSUBSBS, 0xfc0007ff, 0x10000700, 0x0, // Vector Subtract Signed Byte Saturate VX-form (vsubsbs VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSUBSHS, 0xfc0007ff, 0x10000740, 0x0, // Vector Subtract Signed Halfword Saturate VX-form (vsubshs VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSUBSWS, 0xfc0007ff, 0x10000780, 0x0, // Vector Subtract Signed Word Saturate VX-form (vsubsws VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSUBUBM, 0xfc0007ff, 0x10000400, 0x0, // Vector Subtract Unsigned Byte Modulo VX-form (vsububm VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSUBUDM, 0xfc0007ff, 0x100004c0, 0x0, // Vector Subtract Unsigned Doubleword Modulo VX-form (vsubudm VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSUBUHM, 0xfc0007ff, 0x10000440, 0x0, // Vector Subtract Unsigned Halfword Modulo VX-form (vsubuhm VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSUBUWM, 0xfc0007ff, 0x10000480, 0x0, // Vector Subtract Unsigned Word Modulo VX-form (vsubuwm VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSUBUBS, 0xfc0007ff, 0x10000600, 0x0, // Vector Subtract Unsigned Byte Saturate VX-form (vsububs VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSUBUHS, 0xfc0007ff, 0x10000640, 0x0, // Vector Subtract Unsigned Halfword Saturate VX-form (vsubuhs VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSUBUWS, 0xfc0007ff, 0x10000680, 0x0, // Vector Subtract Unsigned Word Saturate VX-form (vsubuws VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSUBUQM, 0xfc0007ff, 0x10000500, 0x0, // Vector Subtract Unsigned Quadword Modulo VX-form (vsubuqm VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSUBEUQM, 0xfc00003f, 0x1000003e, 0x0, // Vector Subtract Extended Unsigned Quadword Modulo VA-form (vsubeuqm VRT,VRA,VRB,VRC)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_VecReg_21_25}},
	{VSUBCUQ, 0xfc0007ff, 0x10000540, 0x0, // Vector Subtract & write Carry Unsigned Quadword VX-form (vsubcuq VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSUBECUQ, 0xfc00003f, 0x1000003f, 0x0, // Vector Subtract Extended & write Carry Unsigned Quadword VA-form (vsubecuq VRT,VRA,VRB,VRC)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_VecReg_21_25}},
	{VMULESB, 0xfc0007ff, 0x10000308, 0x0, // Vector Multiply Even Signed Byte VX-form (vmulesb VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMULEUB, 0xfc0007ff, 0x10000208, 0x0, // Vector Multiply Even Unsigned Byte VX-form (vmuleub VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMULOSB, 0xfc0007ff, 0x10000108, 0x0, // Vector Multiply Odd Signed Byte VX-form (vmulosb VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMULOUB, 0xfc0007ff, 0x10000008, 0x0, // Vector Multiply Odd Unsigned Byte VX-form (vmuloub VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMULESH, 0xfc0007ff, 0x10000348, 0x0, // Vector Multiply Even Signed Halfword VX-form (vmulesh VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMULEUH, 0xfc0007ff, 0x10000248, 0x0, // Vector Multiply Even Unsigned Halfword VX-form (vmuleuh VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMULOSH, 0xfc0007ff, 0x10000148, 0x0, // Vector Multiply Odd Signed Halfword VX-form (vmulosh VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMULOUH, 0xfc0007ff, 0x10000048, 0x0, // Vector Multiply Odd Unsigned Halfword VX-form (vmulouh VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMULESW, 0xfc0007ff, 0x10000388, 0x0, // Vector Multiply Even Signed Word VX-form (vmulesw VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMULEUW, 0xfc0007ff, 0x10000288, 0x0, // Vector Multiply Even Unsigned Word VX-form (vmuleuw VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMULOSW, 0xfc0007ff, 0x10000188, 0x0, // Vector Multiply Odd Signed Word VX-form (vmulosw VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMULOUW, 0xfc0007ff, 0x10000088, 0x0, // Vector Multiply Odd Unsigned Word VX-form (vmulouw VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMULUWM, 0xfc0007ff, 0x10000089, 0x0, // Vector Multiply Unsigned Word Modulo VX-form (vmuluwm VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMHADDSHS, 0xfc00003f, 0x10000020, 0x0, // Vector Multiply-High-Add Signed Halfword Saturate VA-form (vmhaddshs VRT,VRA,VRB,VRC)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_VecReg_21_25}},
	{VMHRADDSHS, 0xfc00003f, 0x10000021, 0x0, // Vector Multiply-High-Round-Add Signed Halfword Saturate VA-form (vmhraddshs VRT,VRA,VRB,VRC)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_VecReg_21_25}},
	{VMLADDUHM, 0xfc00003f, 0x10000022, 0x0, // Vector Multiply-Low-Add Unsigned Halfword Modulo VA-form (vmladduhm VRT,VRA,VRB,VRC)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_VecReg_21_25}},
	{VMSUMUBM, 0xfc00003f, 0x10000024, 0x0, // Vector Multiply-Sum Unsigned Byte Modulo VA-form (vmsumubm VRT,VRA,VRB,VRC)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_VecReg_21_25}},
	{VMSUMMBM, 0xfc00003f, 0x10000025, 0x0, // Vector Multiply-Sum Mixed Byte Modulo VA-form (vmsummbm VRT,VRA,VRB,VRC)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_VecReg_21_25}},
	{VMSUMSHM, 0xfc00003f, 0x10000028, 0x0, // Vector Multiply-Sum Signed Halfword Modulo VA-form (vmsumshm VRT,VRA,VRB,VRC)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_VecReg_21_25}},
	{VMSUMSHS, 0xfc00003f, 0x10000029, 0x0, // Vector Multiply-Sum Signed Halfword Saturate VA-form (vmsumshs VRT,VRA,VRB,VRC)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_VecReg_21_25}},
	{VMSUMUHM, 0xfc00003f, 0x10000026, 0x0, // Vector Multiply-Sum Unsigned Halfword Modulo VA-form (vmsumuhm VRT,VRA,VRB,VRC)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_VecReg_21_25}},
	{VMSUMUHS, 0xfc00003f, 0x10000027, 0x0, // Vector Multiply-Sum Unsigned Halfword Saturate VA-form (vmsumuhs VRT,VRA,VRB,VRC)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_VecReg_21_25}},
	{VSUMSWS, 0xfc0007ff, 0x10000788, 0x0, // Vector Sum across Signed Word Saturate VX-form (vsumsws VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSUM2SWS, 0xfc0007ff, 0x10000688, 0x0, // Vector Sum across Half Signed Word Saturate VX-form (vsum2sws VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSUM4SBS, 0xfc0007ff, 0x10000708, 0x0, // Vector Sum across Quarter Signed Byte Saturate VX-form (vsum4sbs VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSUM4SHS, 0xfc0007ff, 0x10000648, 0x0, // Vector Sum across Quarter Signed Halfword Saturate VX-form (vsum4shs VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSUM4UBS, 0xfc0007ff, 0x10000608, 0x0, // Vector Sum across Quarter Unsigned Byte Saturate VX-form (vsum4ubs VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VAVGSB, 0xfc0007ff, 0x10000502, 0x0, // Vector Average Signed Byte VX-form (vavgsb VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VAVGSH, 0xfc0007ff, 0x10000542, 0x0, // Vector Average Signed Halfword VX-form (vavgsh VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VAVGSW, 0xfc0007ff, 0x10000582, 0x0, // Vector Average Signed Word VX-form (vavgsw VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VAVGUB, 0xfc0007ff, 0x10000402, 0x0, // Vector Average Unsigned Byte VX-form (vavgub VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VAVGUW, 0xfc0007ff, 0x10000482, 0x0, // Vector Average Unsigned Word VX-form (vavguw VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VAVGUH, 0xfc0007ff, 0x10000442, 0x0, // Vector Average Unsigned Halfword VX-form (vavguh VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMAXSB, 0xfc0007ff, 0x10000102, 0x0, // Vector Maximum Signed Byte VX-form (vmaxsb VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMAXSD, 0xfc0007ff, 0x100001c2, 0x0, // Vector Maximum Signed Doubleword VX-form (vmaxsd VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMAXUB, 0xfc0007ff, 0x10000002, 0x0, // Vector Maximum Unsigned Byte VX-form (vmaxub VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMAXUD, 0xfc0007ff, 0x100000c2, 0x0, // Vector Maximum Unsigned Doubleword VX-form (vmaxud VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMAXSH, 0xfc0007ff, 0x10000142, 0x0, // Vector Maximum Signed Halfword VX-form (vmaxsh VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMAXSW, 0xfc0007ff, 0x10000182, 0x0, // Vector Maximum Signed Word VX-form (vmaxsw VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMAXUH, 0xfc0007ff, 0x10000042, 0x0, // Vector Maximum Unsigned Halfword VX-form (vmaxuh VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMAXUW, 0xfc0007ff, 0x10000082, 0x0, // Vector Maximum Unsigned Word VX-form (vmaxuw VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMINSB, 0xfc0007ff, 0x10000302, 0x0, // Vector Minimum Signed Byte VX-form (vminsb VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMINSD, 0xfc0007ff, 0x100003c2, 0x0, // Vector Minimum Signed Doubleword VX-form (vminsd VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMINUB, 0xfc0007ff, 0x10000202, 0x0, // Vector Minimum Unsigned Byte VX-form (vminub VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMINUD, 0xfc0007ff, 0x100002c2, 0x0, // Vector Minimum Unsigned Doubleword VX-form (vminud VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMINSH, 0xfc0007ff, 0x10000342, 0x0, // Vector Minimum Signed Halfword VX-form (vminsh VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMINSW, 0xfc0007ff, 0x10000382, 0x0, // Vector Minimum Signed Word VX-form (vminsw VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMINUH, 0xfc0007ff, 0x10000242, 0x0, // Vector Minimum Unsigned Halfword VX-form (vminuh VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMINUW, 0xfc0007ff, 0x10000282, 0x0, // Vector Minimum Unsigned Word VX-form (vminuw VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPEQUB, 0xfc0007ff, 0x10000006, 0x0, // Vector Compare Equal To Unsigned Byte VC-form (vcmpequb VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPEQUB_, 0xfc0007ff, 0x10000406, 0x0, // Vector Compare Equal To Unsigned Byte VC-form (vcmpequb. VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPEQUH, 0xfc0007ff, 0x10000046, 0x0, // Vector Compare Equal To Unsigned Halfword VC-form (vcmpequh VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPEQUH_, 0xfc0007ff, 0x10000446, 0x0, // Vector Compare Equal To Unsigned Halfword VC-form (vcmpequh. VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPEQUW, 0xfc0007ff, 0x10000086, 0x0, // Vector Compare Equal To Unsigned Word VC-form (vcmpequw VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPEQUW_, 0xfc0007ff, 0x10000486, 0x0, // Vector Compare Equal To Unsigned Word VC-form (vcmpequw. VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPEQUD, 0xfc0007ff, 0x100000c7, 0x0, // Vector Compare Equal To Unsigned Doubleword VX-form (vcmpequd VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPEQUD_, 0xfc0007ff, 0x100004c7, 0x0, // Vector Compare Equal To Unsigned Doubleword VX-form (vcmpequd. VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGTSB, 0xfc0007ff, 0x10000306, 0x0, // Vector Compare Greater Than Signed Byte VC-form (vcmpgtsb VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGTSB_, 0xfc0007ff, 0x10000706, 0x0, // Vector Compare Greater Than Signed Byte VC-form (vcmpgtsb. VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGTSD, 0xfc0007ff, 0x100003c7, 0x0, // Vector Compare Greater Than Signed Doubleword VX-form (vcmpgtsd VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGTSD_, 0xfc0007ff, 0x100007c7, 0x0, // Vector Compare Greater Than Signed Doubleword VX-form (vcmpgtsd. VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGTSH, 0xfc0007ff, 0x10000346, 0x0, // Vector Compare Greater Than Signed Halfword VC-form (vcmpgtsh VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGTSH_, 0xfc0007ff, 0x10000746, 0x0, // Vector Compare Greater Than Signed Halfword VC-form (vcmpgtsh. VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGTSW, 0xfc0007ff, 0x10000386, 0x0, // Vector Compare Greater Than Signed Word VC-form (vcmpgtsw VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGTSW_, 0xfc0007ff, 0x10000786, 0x0, // Vector Compare Greater Than Signed Word VC-form (vcmpgtsw. VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGTUB, 0xfc0007ff, 0x10000206, 0x0, // Vector Compare Greater Than Unsigned Byte VC-form (vcmpgtub VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGTUB_, 0xfc0007ff, 0x10000606, 0x0, // Vector Compare Greater Than Unsigned Byte VC-form (vcmpgtub. VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGTUD, 0xfc0007ff, 0x100002c7, 0x0, // Vector Compare Greater Than Unsigned Doubleword VX-form (vcmpgtud VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGTUD_, 0xfc0007ff, 0x100006c7, 0x0, // Vector Compare Greater Than Unsigned Doubleword VX-form (vcmpgtud. VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGTUH, 0xfc0007ff, 0x10000246, 0x0, // Vector Compare Greater Than Unsigned Halfword VC-form (vcmpgtuh VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGTUH_, 0xfc0007ff, 0x10000646, 0x0, // Vector Compare Greater Than Unsigned Halfword VC-form (vcmpgtuh. VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGTUW, 0xfc0007ff, 0x10000286, 0x0, // Vector Compare Greater Than Unsigned Word VC-form (vcmpgtuw VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGTUW_, 0xfc0007ff, 0x10000686, 0x0, // Vector Compare Greater Than Unsigned Word VC-form (vcmpgtuw. VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VAND, 0xfc0007ff, 0x10000404, 0x0, // Vector Logical AND VX-form (vand VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VANDC, 0xfc0007ff, 0x10000444, 0x0, // Vector Logical AND with Complement VX-form (vandc VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VEQV, 0xfc0007ff, 0x10000684, 0x0, // Vector Logical Equivalent VX-form (veqv VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VNAND, 0xfc0007ff, 0x10000584, 0x0, // Vector Logical NAND VX-form (vnand VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VORC, 0xfc0007ff, 0x10000544, 0x0, // Vector Logical OR with Complement VX-form (vorc VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VNOR, 0xfc0007ff, 0x10000504, 0x0, // Vector Logical NOR VX-form (vnor VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VOR, 0xfc0007ff, 0x10000484, 0x0, // Vector Logical OR VX-form (vor VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VXOR, 0xfc0007ff, 0x100004c4, 0x0, // Vector Logical XOR VX-form (vxor VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VRLB, 0xfc0007ff, 0x10000004, 0x0, // Vector Rotate Left Byte VX-form (vrlb VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VRLH, 0xfc0007ff, 0x10000044, 0x0, // Vector Rotate Left Halfword VX-form (vrlh VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VRLW, 0xfc0007ff, 0x10000084, 0x0, // Vector Rotate Left Word VX-form (vrlw VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VRLD, 0xfc0007ff, 0x100000c4, 0x0, // Vector Rotate Left Doubleword VX-form (vrld VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSLB, 0xfc0007ff, 0x10000104, 0x0, // Vector Shift Left Byte VX-form (vslb VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSLH, 0xfc0007ff, 0x10000144, 0x0, // Vector Shift Left Halfword VX-form (vslh VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSLW, 0xfc0007ff, 0x10000184, 0x0, // Vector Shift Left Word VX-form (vslw VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSLD, 0xfc0007ff, 0x100005c4, 0x0, // Vector Shift Left Doubleword VX-form (vsld VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSRB, 0xfc0007ff, 0x10000204, 0x0, // Vector Shift Right Byte VX-form (vsrb VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSRH, 0xfc0007ff, 0x10000244, 0x0, // Vector Shift Right Halfword VX-form (vsrh VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSRW, 0xfc0007ff, 0x10000284, 0x0, // Vector Shift Right Word VX-form (vsrw VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSRD, 0xfc0007ff, 0x100006c4, 0x0, // Vector Shift Right Doubleword VX-form (vsrd VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSRAB, 0xfc0007ff, 0x10000304, 0x0, // Vector Shift Right Algebraic Byte VX-form (vsrab VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSRAH, 0xfc0007ff, 0x10000344, 0x0, // Vector Shift Right Algebraic Halfword VX-form (vsrah VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSRAW, 0xfc0007ff, 0x10000384, 0x0, // Vector Shift Right Algebraic Word VX-form (vsraw VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSRAD, 0xfc0007ff, 0x100003c4, 0x0, // Vector Shift Right Algebraic Doubleword VX-form (vsrad VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VADDFP, 0xfc0007ff, 0x1000000a, 0x0, // Vector Add Single-Precision VX-form (vaddfp VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSUBFP, 0xfc0007ff, 0x1000004a, 0x0, // Vector Subtract Single-Precision VX-form (vsubfp VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMADDFP, 0xfc00003f, 0x1000002e, 0x0, // Vector Multiply-Add Single-Precision VA-form (vmaddfp VRT,VRA,VRC,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_21_25, ap_VecReg_16_20}},
	{VNMSUBFP, 0xfc00003f, 0x1000002f, 0x0, // Vector Negative Multiply-Subtract Single-Precision VA-form (vnmsubfp VRT,VRA,VRC,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_21_25, ap_VecReg_16_20}},
	{VMAXFP, 0xfc0007ff, 0x1000040a, 0x0, // Vector Maximum Single-Precision VX-form (vmaxfp VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMINFP, 0xfc0007ff, 0x1000044a, 0x0, // Vector Minimum Single-Precision VX-form (vminfp VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCTSXS, 0xfc0007ff, 0x100003ca, 0x0, // Vector Convert To Signed Fixed-Point Word Saturate VX-form (vctsxs VRT,VRB,UIM)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20, ap_ImmUnsigned_11_15}},
	{VCTUXS, 0xfc0007ff, 0x1000038a, 0x0, // Vector Convert To Unsigned Fixed-Point Word Saturate VX-form (vctuxs VRT,VRB,UIM)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20, ap_ImmUnsigned_11_15}},
	{VCFSX, 0xfc0007ff, 0x1000034a, 0x0, // Vector Convert From Signed Fixed-Point Word VX-form (vcfsx VRT,VRB,UIM)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20, ap_ImmUnsigned_11_15}},
	{VCFUX, 0xfc0007ff, 0x1000030a, 0x0, // Vector Convert From Unsigned Fixed-Point Word VX-form (vcfux VRT,VRB,UIM)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20, ap_ImmUnsigned_11_15}},
	{VRFIM, 0xfc0007ff, 0x100002ca, 0x1f0000, // Vector Round to Single-Precision Integer toward -Infinity VX-form (vrfim VRT,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VRFIN, 0xfc0007ff, 0x1000020a, 0x1f0000, // Vector Round to Single-Precision Integer Nearest VX-form (vrfin VRT,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VRFIP, 0xfc0007ff, 0x1000028a, 0x1f0000, // Vector Round to Single-Precision Integer toward +Infinity VX-form (vrfip VRT,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VRFIZ, 0xfc0007ff, 0x1000024a, 0x1f0000, // Vector Round to Single-Precision Integer toward Zero VX-form (vrfiz VRT,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VCMPBFP, 0xfc0007ff, 0x100003c6, 0x0, // Vector Compare Bounds Single-Precision VC-form (vcmpbfp VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPBFP_, 0xfc0007ff, 0x100007c6, 0x0, // Vector Compare Bounds Single-Precision VC-form (vcmpbfp. VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPEQFP, 0xfc0007ff, 0x100000c6, 0x0, // Vector Compare Equal To Single-Precision VC-form (vcmpeqfp VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPEQFP_, 0xfc0007ff, 0x100004c6, 0x0, // Vector Compare Equal To Single-Precision VC-form (vcmpeqfp. VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGEFP, 0xfc0007ff, 0x100001c6, 0x0, // Vector Compare Greater Than or Equal To Single-Precision VC-form (vcmpgefp VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGEFP_, 0xfc0007ff, 0x100005c6, 0x0, // Vector Compare Greater Than or Equal To Single-Precision VC-form (vcmpgefp. VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGTFP, 0xfc0007ff, 0x100002c6, 0x0, // Vector Compare Greater Than Single-Precision VC-form (vcmpgtfp VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGTFP_, 0xfc0007ff, 0x100006c6, 0x0, // Vector Compare Greater Than Single-Precision VC-form (vcmpgtfp. VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VEXPTEFP, 0xfc0007ff, 0x1000018a, 0x1f0000, // Vector 2 Raised to the Exponent Estimate Floating-Point VX-form (vexptefp VRT,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VLOGEFP, 0xfc0007ff, 0x100001ca, 0x1f0000, // Vector Log Base 2 Estimate Floating-Point VX-form (vlogefp VRT,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VREFP, 0xfc0007ff, 0x1000010a, 0x1f0000, // Vector Reciprocal Estimate Single-Precision VX-form (vrefp VRT,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VRSQRTEFP, 0xfc0007ff, 0x1000014a, 0x1f0000, // Vector Reciprocal Square Root Estimate Single-Precision VX-form (vrsqrtefp VRT,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VCIPHER, 0xfc0007ff, 0x10000508, 0x0, // Vector AES Cipher VX-form (vcipher VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCIPHERLAST, 0xfc0007ff, 0x10000509, 0x0, // Vector AES Cipher Last VX-form (vcipherlast VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VNCIPHER, 0xfc0007ff, 0x10000548, 0x0, // Vector AES Inverse Cipher VX-form (vncipher VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VNCIPHERLAST, 0xfc0007ff, 0x10000549, 0x0, // Vector AES Inverse Cipher Last VX-form (vncipherlast VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSBOX, 0xfc0007ff, 0x100005c8, 0xf800, // Vector AES SubBytes VX-form (vsbox VRT,VRA)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15}},
	{VSHASIGMAD, 0xfc0007ff, 0x100006c2, 0x0, // Vector SHA-512 Sigma Doubleword VX-form (vshasigmad VRT,VRA,ST,SIX)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_ImmUnsigned_16_16, ap_ImmUnsigned_17_20}},
	{VSHASIGMAW, 0xfc0007ff, 0x10000682, 0x0, // Vector SHA-256 Sigma Word VX-form (vshasigmaw VRT,VRA,ST,SIX)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_ImmUnsigned_16_16, ap_ImmUnsigned_17_20}},
	{VPMSUMB, 0xfc0007ff, 0x10000408, 0x0, // Vector Polynomial Multiply-Sum Byte VX-form (vpmsumb VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VPMSUMD, 0xfc0007ff, 0x100004c8, 0x0, // Vector Polynomial Multiply-Sum Doubleword VX-form (vpmsumd VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VPMSUMH, 0xfc0007ff, 0x10000448, 0x0, // Vector Polynomial Multiply-Sum Halfword VX-form (vpmsumh VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VPMSUMW, 0xfc0007ff, 0x10000488, 0x0, // Vector Polynomial Multiply-Sum Word VX-form (vpmsumw VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VPERMXOR, 0xfc00003f, 0x1000002d, 0x0, // Vector Permute and Exclusive-OR VA-form (vpermxor VRT,VRA,VRB,VRC)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_VecReg_21_25}},
	{VGBBD, 0xfc0007ff, 0x1000050c, 0x1f0000, // Vector Gather Bits by Bytes by Doubleword VX-form (vgbbd VRT,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VCLZB, 0xfc0007ff, 0x10000702, 0x1f0000, // Vector Count Leading Zeros Byte VX-form (vclzb VRT,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VCLZH, 0xfc0007ff, 0x10000742, 0x1f0000, // Vector Count Leading Zeros Halfword VX-form (vclzh VRT,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VCLZW, 0xfc0007ff, 0x10000782, 0x1f0000, // Vector Count Leading Zeros Word VX-form (vclzw VRT,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VCLZD, 0xfc0007ff, 0x100007c2, 0x1f0000, // Vector Count Leading Zeros Doubleword (vclzd VRT,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VPOPCNTB, 0xfc0007ff, 0x10000703, 0x1f0000, // Vector Population Count Byte (vpopcntb VRT,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VPOPCNTD, 0xfc0007ff, 0x100007c3, 0x1f0000, // Vector Population Count Doubleword (vpopcntd VRT,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VPOPCNTH, 0xfc0007ff, 0x10000743, 0x1f0000, // Vector Population Count Halfword (vpopcnth VRT,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VPOPCNTW, 0xfc0007ff, 0x10000783, 0x1f0000, // Vector Population Count Word (vpopcntw VRT,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VBPERMQ, 0xfc0007ff, 0x1000054c, 0x0, // Vector Bit Permute Quadword VX-form (vbpermq VRT,VRA,VRB)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{BCDADD_, 0xfc0005ff, 0x10000401, 0x0, // Decimal Add Modulo VX-form (bcdadd. VRT,VRA,VRB,PS)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_ImmUnsigned_22_22}},
	{BCDSUB_, 0xfc0005ff, 0x10000441, 0x0, // Decimal Subtract Modulo VX-form (bcdsub. VRT,VRA,VRB,PS)
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_ImmUnsigned_22_22}},
	{MTVSCR, 0xfc0007ff, 0x10000644, 0x3ff0000, // Move To Vector Status and Control Register VX-form (mtvscr VRB)
		[5]*argField{ap_VecReg_16_20}},
	{MFVSCR, 0xfc0007ff, 0x10000604, 0x1ff800, // Move From Vector Status and Control Register VX-form (mfvscr VRT)
		[5]*argField{ap_VecReg_6_10}},
	{DADD, 0xfc0007ff, 0xec000004, 0x0, // DFP Add [Quad] X-form (dadd FRT,FRA,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{DADD_, 0xfc0007ff, 0xec000005, 0x0, // DFP Add [Quad] X-form (dadd. FRT,FRA,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{DSUB, 0xfc0007ff, 0xec000404, 0x0, // DFP Subtract [Quad] X-form (dsub FRT,FRA,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{DSUB_, 0xfc0007ff, 0xec000405, 0x0, // DFP Subtract [Quad] X-form (dsub. FRT,FRA,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{DMUL, 0xfc0007ff, 0xec000044, 0x0, // DFP Multiply [Quad] X-form (dmul FRT,FRA,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{DMUL_, 0xfc0007ff, 0xec000045, 0x0, // DFP Multiply [Quad] X-form (dmul. FRT,FRA,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{DDIV, 0xfc0007ff, 0xec000444, 0x0, // DFP Divide [Quad] X-form (ddiv FRT,FRA,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{DDIV_, 0xfc0007ff, 0xec000445, 0x0, // DFP Divide [Quad] X-form (ddiv. FRT,FRA,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{DCMPU, 0xfc0007fe, 0xec000504, 0x600001, // DFP Compare Unordered [Quad] X-form (dcmpu BF,FRA,FRB)
		[5]*argField{ap_CondRegField_6_8, ap_FPReg_11_15, ap_FPReg_16_20}},
	{DCMPO, 0xfc0007fe, 0xec000104, 0x600001, // DFP Compare Ordered [Quad] X-form (dcmpo BF,FRA,FRB)
		[5]*argField{ap_CondRegField_6_8, ap_FPReg_11_15, ap_FPReg_16_20}},
	{DTSTDC, 0xfc0003fe, 0xec000184, 0x600001, // DFP Test Data Class [Quad] Z22-form (dtstdc BF,FRA,DCM)
		[5]*argField{ap_CondRegField_6_8, ap_FPReg_11_15, ap_ImmUnsigned_16_21}},
	{DTSTDG, 0xfc0003fe, 0xec0001c4, 0x600001, // DFP Test Data Group [Quad] Z22-form (dtstdg BF,FRA,DGM)
		[5]*argField{ap_CondRegField_6_8, ap_FPReg_11_15, ap_ImmUnsigned_16_21}},
	{DTSTEX, 0xfc0007fe, 0xec000144, 0x600001, // DFP Test Exponent [Quad] X-form (dtstex BF,FRA,FRB)
		[5]*argField{ap_CondRegField_6_8, ap_FPReg_11_15, ap_FPReg_16_20}},
	{DTSTSF, 0xfc0007fe, 0xec000544, 0x600001, // DFP Test Significance [Quad] X-form (dtstsf BF,FRA,FRB)
		[5]*argField{ap_CondRegField_6_8, ap_FPReg_11_15, ap_FPReg_16_20}},
	{DQUAI, 0xfc0001ff, 0xec000086, 0x0, // DFP Quantize Immediate [Quad] Z23-form (dquai TE,FRT,FRB,RMC)
		[5]*argField{ap_ImmSigned_11_15, ap_FPReg_6_10, ap_FPReg_16_20, ap_ImmUnsigned_21_22}},
	{DQUAI_, 0xfc0001ff, 0xec000087, 0x0, // DFP Quantize Immediate [Quad] Z23-form (dquai. TE,FRT,FRB,RMC)
		[5]*argField{ap_ImmSigned_11_15, ap_FPReg_6_10, ap_FPReg_16_20, ap_ImmUnsigned_21_22}},
	{DQUA, 0xfc0001ff, 0xec000006, 0x0, // DFP Quantize [Quad] Z23-form (dqua FRT,FRA,FRB,RMC)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20, ap_ImmUnsigned_21_22}},
	{DQUA_, 0xfc0001ff, 0xec000007, 0x0, // DFP Quantize [Quad] Z23-form (dqua. FRT,FRA,FRB,RMC)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20, ap_ImmUnsigned_21_22}},
	{DRRND, 0xfc0001ff, 0xec000046, 0x0, // DFP Reround [Quad] Z23-form (drrnd FRT,FRA,FRB,RMC)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20, ap_ImmUnsigned_21_22}},
	{DRRND_, 0xfc0001ff, 0xec000047, 0x0, // DFP Reround [Quad] Z23-form (drrnd. FRT,FRA,FRB,RMC)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20, ap_ImmUnsigned_21_22}},
	{DRINTX, 0xfc0001ff, 0xec0000c6, 0x1e0000, // DFP Round To FP Integer With Inexact [Quad] Z23-form (drintx R,FRT,FRB,RMC)
		[5]*argField{ap_ImmUnsigned_15_15, ap_FPReg_6_10, ap_FPReg_16_20, ap_ImmUnsigned_21_22}},
	{DRINTX_, 0xfc0001ff, 0xec0000c7, 0x1e0000, // DFP Round To FP Integer With Inexact [Quad] Z23-form (drintx. R,FRT,FRB,RMC)
		[5]*argField{ap_ImmUnsigned_15_15, ap_FPReg_6_10, ap_FPReg_16_20, ap_ImmUnsigned_21_22}},
	{DRINTN, 0xfc0001ff, 0xec0001c6, 0x1e0000, // DFP Round To FP Integer Without Inexact [Quad] Z23-form (drintn R,FRT,FRB,RMC)
		[5]*argField{ap_ImmUnsigned_15_15, ap_FPReg_6_10, ap_FPReg_16_20, ap_ImmUnsigned_21_22}},
	{DRINTN_, 0xfc0001ff, 0xec0001c7, 0x1e0000, // DFP Round To FP Integer Without Inexact [Quad] Z23-form (drintn. R,FRT,FRB,RMC)
		[5]*argField{ap_ImmUnsigned_15_15, ap_FPReg_6_10, ap_FPReg_16_20, ap_ImmUnsigned_21_22}},
	{DCTDP, 0xfc0007ff, 0xec000204, 0x1f0000, // DFP Convert To DFP Long X-form (dctdp FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{DCTDP_, 0xfc0007ff, 0xec000205, 0x1f0000, // DFP Convert To DFP Long X-form (dctdp. FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{DCTQPQ, 0xfc0007ff, 0xfc000204, 0x1f0000, // DFP Convert To DFP Extended X-form (dctqpq FRTp,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{DCTQPQ_, 0xfc0007ff, 0xfc000205, 0x1f0000, // DFP Convert To DFP Extended X-form (dctqpq. FRTp,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{DRSP, 0xfc0007ff, 0xec000604, 0x1f0000, // DFP Round To DFP Short X-form (drsp FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{DRSP_, 0xfc0007ff, 0xec000605, 0x1f0000, // DFP Round To DFP Short X-form (drsp. FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{DRDPQ, 0xfc0007ff, 0xfc000604, 0x1f0000, // DFP Round To DFP Long X-form (drdpq FRTp,FRBp)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{DRDPQ_, 0xfc0007ff, 0xfc000605, 0x1f0000, // DFP Round To DFP Long X-form (drdpq. FRTp,FRBp)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{DCFFIX, 0xfc0007ff, 0xec000644, 0x1f0000, // DFP Convert From Fixed X-form (dcffix FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{DCFFIX_, 0xfc0007ff, 0xec000645, 0x1f0000, // DFP Convert From Fixed X-form (dcffix. FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{DCFFIXQ, 0xfc0007ff, 0xfc000644, 0x1f0000, // DFP Convert From Fixed Quad X-form (dcffixq FRTp,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{DCFFIXQ_, 0xfc0007ff, 0xfc000645, 0x1f0000, // DFP Convert From Fixed Quad X-form (dcffixq. FRTp,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{DCTFIX, 0xfc0007ff, 0xec000244, 0x1f0000, // DFP Convert To Fixed [Quad] X-form (dctfix FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{DCTFIX_, 0xfc0007ff, 0xec000245, 0x1f0000, // DFP Convert To Fixed [Quad] X-form (dctfix. FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{DDEDPD, 0xfc0007ff, 0xec000284, 0x70000, // DFP Decode DPD To BCD [Quad] X-form (ddedpd SP,FRT,FRB)
		[5]*argField{ap_ImmUnsigned_11_12, ap_FPReg_6_10, ap_FPReg_16_20}},
	{DDEDPD_, 0xfc0007ff, 0xec000285, 0x70000, // DFP Decode DPD To BCD [Quad] X-form (ddedpd. SP,FRT,FRB)
		[5]*argField{ap_ImmUnsigned_11_12, ap_FPReg_6_10, ap_FPReg_16_20}},
	{DENBCD, 0xfc0007ff, 0xec000684, 0xf0000, // DFP Encode BCD To DPD [Quad] X-form (denbcd S,FRT,FRB)
		[5]*argField{ap_ImmUnsigned_11_11, ap_FPReg_6_10, ap_FPReg_16_20}},
	{DENBCD_, 0xfc0007ff, 0xec000685, 0xf0000, // DFP Encode BCD To DPD [Quad] X-form (denbcd. S,FRT,FRB)
		[5]*argField{ap_ImmUnsigned_11_11, ap_FPReg_6_10, ap_FPReg_16_20}},
	{DXEX, 0xfc0007ff, 0xec0002c4, 0x1f0000, // DFP Extract Biased Exponent [Quad] X-form (dxex FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{DXEX_, 0xfc0007ff, 0xec0002c5, 0x1f0000, // DFP Extract Biased Exponent [Quad] X-form (dxex. FRT,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{DIEX, 0xfc0007ff, 0xec0006c4, 0x0, // DFP Insert Biased Exponent [Quad] X-form (diex FRT,FRA,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{DIEX_, 0xfc0007ff, 0xec0006c5, 0x0, // DFP Insert Biased Exponent [Quad] X-form (diex. FRT,FRA,FRB)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{DSCLI, 0xfc0003ff, 0xec000084, 0x0, // DFP Shift Significand Left Immediate [Quad] Z22-form (dscli FRT,FRA,SH)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_ImmUnsigned_16_21}},
	{DSCLI_, 0xfc0003ff, 0xec000085, 0x0, // DFP Shift Significand Left Immediate [Quad] Z22-form (dscli. FRT,FRA,SH)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_ImmUnsigned_16_21}},
	{DSCRI, 0xfc0003ff, 0xec0000c4, 0x0, // DFP Shift Significand Right Immediate [Quad] Z22-form (dscri FRT,FRA,SH)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_ImmUnsigned_16_21}},
	{DSCRI_, 0xfc0003ff, 0xec0000c5, 0x0, // DFP Shift Significand Right Immediate [Quad] Z22-form (dscri. FRT,FRA,SH)
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_ImmUnsigned_16_21}},
	{LXSDX, 0xfc0007fe, 0x7c000498, 0x0, // Load VSX Scalar Doubleword Indexed XX1-form (lxsdx XT,RA,RB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LXSIWAX, 0xfc0007fe, 0x7c000098, 0x0, // Load VSX Scalar as Integer Word Algebraic Indexed XX1-form (lxsiwax XT,RA,RB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LXSIWZX, 0xfc0007fe, 0x7c000018, 0x0, // Load VSX Scalar as Integer Word and Zero Indexed XX1-form (lxsiwzx XT,RA,RB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LXSSPX, 0xfc0007fe, 0x7c000418, 0x0, // Load VSX Scalar Single-Precision Indexed XX1-form (lxsspx XT,RA,RB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LXVD2X, 0xfc0007fe, 0x7c000698, 0x0, // Load VSX Vector Doubleword*2 Indexed XX1-form (lxvd2x XT,RA,RB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LXVDSX, 0xfc0007fe, 0x7c000298, 0x0, // Load VSX Vector Doubleword & Splat Indexed XX1-form (lxvdsx XT,RA,RB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LXVW4X, 0xfc0007fe, 0x7c000618, 0x0, // Load VSX Vector Word*4 Indexed XX1-form (lxvw4x XT,RA,RB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STXSDX, 0xfc0007fe, 0x7c000598, 0x0, // Store VSX Scalar Doubleword Indexed XX1-form (stxsdx XS,RA,RB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STXSIWX, 0xfc0007fe, 0x7c000118, 0x0, // Store VSX Scalar as Integer Word Indexed XX1-form (stxsiwx XS,RA,RB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STXSSPX, 0xfc0007fe, 0x7c000518, 0x0, // Store VSX Scalar Single-Precision Indexed XX1-form (stxsspx XS,RA,RB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STXVD2X, 0xfc0007fe, 0x7c000798, 0x0, // Store VSX Vector Doubleword*2 Indexed XX1-form (stxvd2x XS,RA,RB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STXVW4X, 0xfc0007fe, 0x7c000718, 0x0, // Store VSX Vector Word*4 Indexed XX1-form (stxvw4x XS,RA,RB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{XSABSDP, 0xfc0007fc, 0xf0000564, 0x1f0000, // VSX Scalar Absolute Value Double-Precision XX2-form (xsabsdp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XSADDDP, 0xfc0007f8, 0xf0000100, 0x0, // VSX Scalar Add Double-Precision XX3-form (xsadddp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XSADDSP, 0xfc0007f8, 0xf0000000, 0x0, // VSX Scalar Add Single-Precision XX3-form (xsaddsp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XSCMPODP, 0xfc0007f8, 0xf0000158, 0x600001, // VSX Scalar Compare Ordered Double-Precision XX3-form (xscmpodp BF,XA,XB)
		[5]*argField{ap_CondRegField_6_8, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XSCMPUDP, 0xfc0007f8, 0xf0000118, 0x600001, // VSX Scalar Compare Unordered Double-Precision XX3-form (xscmpudp BF,XA,XB)
		[5]*argField{ap_CondRegField_6_8, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XSCPSGNDP, 0xfc0007f8, 0xf0000580, 0x0, // VSX Scalar Copy Sign Double-Precision XX3-form (xscpsgndp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XSCVDPSP, 0xfc0007fc, 0xf0000424, 0x1f0000, // VSX Scalar round Double-Precision to single-precision and Convert to Single-Precision format XX2-form (xscvdpsp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XSCVDPSPN, 0xfc0007fc, 0xf000042c, 0x1f0000, // VSX Scalar Convert Scalar Single-Precision to Vector Single-Precision format Non-signalling XX2-form (xscvdpspn XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XSCVDPSXDS, 0xfc0007fc, 0xf0000560, 0x1f0000, // VSX Scalar truncate Double-Precision to integer and Convert to Signed Integer Doubleword format with Saturate XX2-form (xscvdpsxds XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XSCVDPSXWS, 0xfc0007fc, 0xf0000160, 0x1f0000, // VSX Scalar truncate Double-Precision to integer and Convert to Signed Integer Word format with Saturate XX2-form (xscvdpsxws XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XSCVDPUXDS, 0xfc0007fc, 0xf0000520, 0x1f0000, // VSX Scalar truncate Double-Precision integer and Convert to Unsigned Integer Doubleword format with Saturate XX2-form (xscvdpuxds XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XSCVDPUXWS, 0xfc0007fc, 0xf0000120, 0x1f0000, // VSX Scalar truncate Double-Precision to integer and Convert to Unsigned Integer Word format with Saturate XX2-form (xscvdpuxws XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XSCVSPDP, 0xfc0007fc, 0xf0000524, 0x1f0000, // VSX Scalar Convert Single-Precision to Double-Precision format XX2-form (xscvspdp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XSCVSPDPN, 0xfc0007fc, 0xf000052c, 0x1f0000, // VSX Scalar Convert Single-Precision to Double-Precision format Non-signalling XX2-form (xscvspdpn XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XSCVSXDDP, 0xfc0007fc, 0xf00005e0, 0x1f0000, // VSX Scalar Convert Signed Integer Doubleword to floating-point format and round to Double-Precision format XX2-form (xscvsxddp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XSCVSXDSP, 0xfc0007fc, 0xf00004e0, 0x1f0000, // VSX Scalar Convert Signed Integer Doubleword to floating-point format and round to Single-Precision XX2-form (xscvsxdsp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XSCVUXDDP, 0xfc0007fc, 0xf00005a0, 0x1f0000, // VSX Scalar Convert Unsigned Integer Doubleword to floating-point format and round to Double-Precision format XX2-form (xscvuxddp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XSCVUXDSP, 0xfc0007fc, 0xf00004a0, 0x1f0000, // VSX Scalar Convert Unsigned Integer Doubleword to floating-point format and round to Single-Precision XX2-form (xscvuxdsp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XSDIVDP, 0xfc0007f8, 0xf00001c0, 0x0, // VSX Scalar Divide Double-Precision XX3-form (xsdivdp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XSDIVSP, 0xfc0007f8, 0xf00000c0, 0x0, // VSX Scalar Divide Single-Precision XX3-form (xsdivsp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XSMADDADP, 0xfc0007f8, 0xf0000108, 0x0, // VSX Scalar Multiply-Add Double-Precision XX3-form (xsmaddadp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XSMADDASP, 0xfc0007f8, 0xf0000008, 0x0, // VSX Scalar Multiply-Add Single-Precision XX3-form (xsmaddasp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XSMAXDP, 0xfc0007f8, 0xf0000500, 0x0, // VSX Scalar Maximum Double-Precision XX3-form (xsmaxdp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XSMINDP, 0xfc0007f8, 0xf0000540, 0x0, // VSX Scalar Minimum Double-Precision XX3-form (xsmindp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XSMSUBADP, 0xfc0007f8, 0xf0000188, 0x0, // VSX Scalar Multiply-Subtract Double-Precision XX3-form (xsmsubadp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XSMSUBASP, 0xfc0007f8, 0xf0000088, 0x0, // VSX Scalar Multiply-Subtract Single-Precision XX3-form (xsmsubasp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XSMULDP, 0xfc0007f8, 0xf0000180, 0x0, // VSX Scalar Multiply Double-Precision XX3-form (xsmuldp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XSMULSP, 0xfc0007f8, 0xf0000080, 0x0, // VSX Scalar Multiply Single-Precision XX3-form (xsmulsp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XSNABSDP, 0xfc0007fc, 0xf00005a4, 0x1f0000, // VSX Scalar Negative Absolute Value Double-Precision XX2-form (xsnabsdp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XSNEGDP, 0xfc0007fc, 0xf00005e4, 0x1f0000, // VSX Scalar Negate Double-Precision XX2-form (xsnegdp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XSNMADDADP, 0xfc0007f8, 0xf0000508, 0x0, // VSX Scalar Negative Multiply-Add Double-Precision XX3-form (xsnmaddadp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XSNMADDASP, 0xfc0007f8, 0xf0000408, 0x0, // VSX Scalar Negative Multiply-Add Single-Precision XX3-form (xsnmaddasp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XSNMSUBADP, 0xfc0007f8, 0xf0000588, 0x0, // VSX Scalar Negative Multiply-Subtract Double-Precision XX3-form (xsnmsubadp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XSNMSUBASP, 0xfc0007f8, 0xf0000488, 0x0, // VSX Scalar Negative Multiply-Subtract Single-Precision XX3-form (xsnmsubasp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XSRDPI, 0xfc0007fc, 0xf0000124, 0x1f0000, // VSX Scalar Round to Double-Precision Integer using round to Nearest Away XX2-form (xsrdpi XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XSRDPIC, 0xfc0007fc, 0xf00001ac, 0x1f0000, // VSX Scalar Round to Double-Precision Integer exact using Current rounding mode XX2-form (xsrdpic XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XSRDPIM, 0xfc0007fc, 0xf00001e4, 0x1f0000, // VSX Scalar Round to Double-Precision Integer using round toward -Infinity XX2-form (xsrdpim XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XSRDPIP, 0xfc0007fc, 0xf00001a4, 0x1f0000, // VSX Scalar Round to Double-Precision Integer using round toward +Infinity XX2-form (xsrdpip XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XSRDPIZ, 0xfc0007fc, 0xf0000164, 0x1f0000, // VSX Scalar Round to Double-Precision Integer using round toward Zero XX2-form (xsrdpiz XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XSREDP, 0xfc0007fc, 0xf0000168, 0x1f0000, // VSX Scalar Reciprocal Estimate Double-Precision XX2-form (xsredp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XSRESP, 0xfc0007fc, 0xf0000068, 0x1f0000, // VSX Scalar Reciprocal Estimate Single-Precision XX2-form (xsresp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XSRSP, 0xfc0007fc, 0xf0000464, 0x1f0000, // VSX Scalar Round to Single-Precision XX2-form (xsrsp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XSRSQRTEDP, 0xfc0007fc, 0xf0000128, 0x1f0000, // VSX Scalar Reciprocal Square Root Estimate Double-Precision XX2-form (xsrsqrtedp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XSRSQRTESP, 0xfc0007fc, 0xf0000028, 0x1f0000, // VSX Scalar Reciprocal Square Root Estimate Single-Precision XX2-form (xsrsqrtesp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XSSQRTDP, 0xfc0007fc, 0xf000012c, 0x1f0000, // VSX Scalar Square Root Double-Precision XX2-form (xssqrtdp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XSSQRTSP, 0xfc0007fc, 0xf000002c, 0x1f0000, // VSX Scalar Square Root Single-Precision XX-form (xssqrtsp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XSSUBDP, 0xfc0007f8, 0xf0000140, 0x0, // VSX Scalar Subtract Double-Precision XX3-form (xssubdp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XSSUBSP, 0xfc0007f8, 0xf0000040, 0x0, // VSX Scalar Subtract Single-Precision XX3-form (xssubsp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XSTDIVDP, 0xfc0007f8, 0xf00001e8, 0x600001, // VSX Scalar Test for software Divide Double-Precision XX3-form (xstdivdp BF,XA,XB)
		[5]*argField{ap_CondRegField_6_8, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XSTSQRTDP, 0xfc0007fc, 0xf00001a8, 0x7f0001, // VSX Scalar Test for software Square Root Double-Precision XX2-form (xstsqrtdp BF,XB)
		[5]*argField{ap_CondRegField_6_8, ap_VecReg_30_30_16_20}},
	{XVABSDP, 0xfc0007fc, 0xf0000764, 0x1f0000, // VSX Vector Absolute Value Double-Precision XX2-form (xvabsdp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVABSSP, 0xfc0007fc, 0xf0000664, 0x1f0000, // VSX Vector Absolute Value Single-Precision XX2-form (xvabssp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVADDDP, 0xfc0007f8, 0xf0000300, 0x0, // VSX Vector Add Double-Precision XX3-form (xvadddp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XVADDSP, 0xfc0007f8, 0xf0000200, 0x0, // VSX Vector Add Single-Precision XX3-form (xvaddsp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XVCMPEQDP, 0xfc0007f8, 0xf0000318, 0x0, // VSX Vector Compare Equal To Double-Precision [ & Record ] XX3-form (xvcmpeqdp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XVCMPEQDP_, 0xfc0007f8, 0xf0000718, 0x0, // VSX Vector Compare Equal To Double-Precision [ & Record ] XX3-form (xvcmpeqdp. XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XVCMPEQSP, 0xfc0007f8, 0xf0000218, 0x0, // VSX Vector Compare Equal To Single-Precision [ & Record ] XX3-form (xvcmpeqsp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XVCMPEQSP_, 0xfc0007f8, 0xf0000618, 0x0, // VSX Vector Compare Equal To Single-Precision [ & Record ] XX3-form (xvcmpeqsp. XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XVCMPGEDP, 0xfc0007f8, 0xf0000398, 0x0, // VSX Vector Compare Greater Than or Equal To Double-Precision [ & Record ] XX3-form (xvcmpgedp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XVCMPGEDP_, 0xfc0007f8, 0xf0000798, 0x0, // VSX Vector Compare Greater Than or Equal To Double-Precision [ & Record ] XX3-form (xvcmpgedp. XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XVCMPGESP, 0xfc0007f8, 0xf0000298, 0x0, // VSX Vector Compare Greater Than or Equal To Single-Precision [ & record CR6 ] XX3-form (xvcmpgesp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XVCMPGESP_, 0xfc0007f8, 0xf0000698, 0x0, // VSX Vector Compare Greater Than or Equal To Single-Precision [ & record CR6 ] XX3-form (xvcmpgesp. XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XVCMPGTDP, 0xfc0007f8, 0xf0000358, 0x0, // VSX Vector Compare Greater Than Double-Precision [ & record CR6 ] XX3-form (xvcmpgtdp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XVCMPGTDP_, 0xfc0007f8, 0xf0000758, 0x0, // VSX Vector Compare Greater Than Double-Precision [ & record CR6 ] XX3-form (xvcmpgtdp. XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XVCMPGTSP, 0xfc0007f8, 0xf0000258, 0x0, // VSX Vector Compare Greater Than Single-Precision [ & record CR6 ] XX3-form (xvcmpgtsp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XVCMPGTSP_, 0xfc0007f8, 0xf0000658, 0x0, // VSX Vector Compare Greater Than Single-Precision [ & record CR6 ] XX3-form (xvcmpgtsp. XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XVCPSGNDP, 0xfc0007f8, 0xf0000780, 0x0, // VSX Vector Copy Sign Double-Precision XX3-form (xvcpsgndp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XVCPSGNSP, 0xfc0007f8, 0xf0000680, 0x0, // VSX Vector Copy Sign Single-Precision XX3-form (xvcpsgnsp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XVCVDPSP, 0xfc0007fc, 0xf0000624, 0x1f0000, // VSX Vector round Double-Precision to single-precision and Convert to Single-Precision format XX2-form (xvcvdpsp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVCVDPSXDS, 0xfc0007fc, 0xf0000760, 0x1f0000, // VSX Vector truncate Double-Precision to integer and Convert to Signed Integer Doubleword format with Saturate XX2-form (xvcvdpsxds XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVCVDPSXWS, 0xfc0007fc, 0xf0000360, 0x1f0000, // VSX Vector truncate Double-Precision to integer and Convert to Signed Integer Word format with Saturate XX2-form (xvcvdpsxws XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVCVDPUXDS, 0xfc0007fc, 0xf0000720, 0x1f0000, // VSX Vector truncate Double-Precision to integer and Convert to Unsigned Integer Doubleword format with Saturate XX2-form (xvcvdpuxds XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVCVDPUXWS, 0xfc0007fc, 0xf0000320, 0x1f0000, // VSX Vector truncate Double-Precision to integer and Convert to Unsigned Integer Word format with Saturate XX2-form (xvcvdpuxws XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVCVSPDP, 0xfc0007fc, 0xf0000724, 0x1f0000, // VSX Vector Convert Single-Precision to Double-Precision format XX2-form (xvcvspdp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVCVSPSXDS, 0xfc0007fc, 0xf0000660, 0x1f0000, // VSX Vector truncate Single-Precision to integer and Convert to Signed Integer Doubleword format with Saturate XX2-form (xvcvspsxds XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVCVSPSXWS, 0xfc0007fc, 0xf0000260, 0x1f0000, // VSX Vector truncate Single-Precision to integer and Convert to Signed Integer Word format with Saturate XX2-form (xvcvspsxws XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVCVSPUXDS, 0xfc0007fc, 0xf0000620, 0x1f0000, // VSX Vector truncate Single-Precision to integer and Convert to Unsigned Integer Doubleword format with Saturate XX2-form (xvcvspuxds XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVCVSPUXWS, 0xfc0007fc, 0xf0000220, 0x1f0000, // VSX Vector truncate Single-Precision to integer and Convert to Unsigned Integer Word format with Saturate XX2-form (xvcvspuxws XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVCVSXDDP, 0xfc0007fc, 0xf00007e0, 0x1f0000, // VSX Vector Convert and round Signed Integer Doubleword to Double-Precision format XX2-form (xvcvsxddp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVCVSXDSP, 0xfc0007fc, 0xf00006e0, 0x1f0000, // VSX Vector Convert and round Signed Integer Doubleword to Single-Precision format XX2-form (xvcvsxdsp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVCVSXWDP, 0xfc0007fc, 0xf00003e0, 0x1f0000, // VSX Vector Convert Signed Integer Word to Double-Precision format XX2-form (xvcvsxwdp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVCVSXWSP, 0xfc0007fc, 0xf00002e0, 0x1f0000, // VSX Vector Convert and round Signed Integer Word to Single-Precision format XX2-form (xvcvsxwsp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVCVUXDDP, 0xfc0007fc, 0xf00007a0, 0x1f0000, // VSX Vector Convert and round Unsigned Integer Doubleword to Double-Precision format XX2-form (xvcvuxddp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVCVUXDSP, 0xfc0007fc, 0xf00006a0, 0x1f0000, // VSX Vector Convert and round Unsigned Integer Doubleword to Single-Precision format XX2-form (xvcvuxdsp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVCVUXWDP, 0xfc0007fc, 0xf00003a0, 0x1f0000, // VSX Vector Convert and round Unsigned Integer Word to Double-Precision format XX2-form (xvcvuxwdp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVCVUXWSP, 0xfc0007fc, 0xf00002a0, 0x1f0000, // VSX Vector Convert and round Unsigned Integer Word to Single-Precision format XX2-form (xvcvuxwsp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVDIVDP, 0xfc0007f8, 0xf00003c0, 0x0, // VSX Vector Divide Double-Precision XX3-form (xvdivdp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XVDIVSP, 0xfc0007f8, 0xf00002c0, 0x0, // VSX Vector Divide Single-Precision XX3-form (xvdivsp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XVMADDADP, 0xfc0007f8, 0xf0000308, 0x0, // VSX Vector Multiply-Add Double-Precision XX3-form (xvmaddadp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XVMADDASP, 0xfc0007f8, 0xf0000208, 0x0, // VSX Vector Multiply-Add Single-Precision XX3-form (xvmaddasp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XVMAXDP, 0xfc0007f8, 0xf0000700, 0x0, // VSX Vector Maximum Double-Precision XX3-form (xvmaxdp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XVMAXSP, 0xfc0007f8, 0xf0000600, 0x0, // VSX Vector Maximum Single-Precision XX3-form (xvmaxsp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XVMINDP, 0xfc0007f8, 0xf0000740, 0x0, // VSX Vector Minimum Double-Precision XX3-form (xvmindp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XVMINSP, 0xfc0007f8, 0xf0000640, 0x0, // VSX Vector Minimum Single-Precision XX3-form (xvminsp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XVMSUBADP, 0xfc0007f8, 0xf0000388, 0x0, // VSX Vector Multiply-Subtract Double-Precision XX3-form (xvmsubadp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XVMSUBASP, 0xfc0007f8, 0xf0000288, 0x0, // VSX Vector Multiply-Subtract Single-Precision XX3-form (xvmsubasp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XVMULDP, 0xfc0007f8, 0xf0000380, 0x0, // VSX Vector Multiply Double-Precision XX3-form (xvmuldp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XVMULSP, 0xfc0007f8, 0xf0000280, 0x0, // VSX Vector Multiply Single-Precision XX3-form (xvmulsp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XVNABSDP, 0xfc0007fc, 0xf00007a4, 0x1f0000, // VSX Vector Negative Absolute Value Double-Precision XX2-form (xvnabsdp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVNABSSP, 0xfc0007fc, 0xf00006a4, 0x1f0000, // VSX Vector Negative Absolute Value Single-Precision XX2-form (xvnabssp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVNEGDP, 0xfc0007fc, 0xf00007e4, 0x1f0000, // VSX Vector Negate Double-Precision XX2-form (xvnegdp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVNEGSP, 0xfc0007fc, 0xf00006e4, 0x1f0000, // VSX Vector Negate Single-Precision XX2-form (xvnegsp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVNMADDADP, 0xfc0007f8, 0xf0000708, 0x0, // VSX Vector Negative Multiply-Add Double-Precision XX3-form (xvnmaddadp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XVNMADDASP, 0xfc0007f8, 0xf0000608, 0x0, // VSX Vector Negative Multiply-Add Single-Precision XX3-form (xvnmaddasp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XVNMSUBADP, 0xfc0007f8, 0xf0000788, 0x0, // VSX Vector Negative Multiply-Subtract Double-Precision XX3-form (xvnmsubadp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XVNMSUBASP, 0xfc0007f8, 0xf0000688, 0x0, // VSX Vector Negative Multiply-Subtract Single-Precision XX3-form (xvnmsubasp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XVRDPI, 0xfc0007fc, 0xf0000324, 0x1f0000, // VSX Vector Round to Double-Precision Integer using round to Nearest Away XX2-form (xvrdpi XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVRDPIC, 0xfc0007fc, 0xf00003ac, 0x1f0000, // VSX Vector Round to Double-Precision Integer Exact using Current rounding mode XX2-form (xvrdpic XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVRDPIM, 0xfc0007fc, 0xf00003e4, 0x1f0000, // VSX Vector Round to Double-Precision Integer using round toward -Infinity XX2-form (xvrdpim XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVRDPIP, 0xfc0007fc, 0xf00003a4, 0x1f0000, // VSX Vector Round to Double-Precision Integer using round toward +Infinity XX2-form (xvrdpip XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVRDPIZ, 0xfc0007fc, 0xf0000364, 0x1f0000, // VSX Vector Round to Double-Precision Integer using round toward Zero XX2-form (xvrdpiz XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVREDP, 0xfc0007fc, 0xf0000368, 0x1f0000, // VSX Vector Reciprocal Estimate Double-Precision XX2-form (xvredp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVRESP, 0xfc0007fc, 0xf0000268, 0x1f0000, // VSX Vector Reciprocal Estimate Single-Precision XX2-form (xvresp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVRSPI, 0xfc0007fc, 0xf0000224, 0x1f0000, // VSX Vector Round to Single-Precision Integer using round to Nearest Away XX2-form (xvrspi XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVRSPIC, 0xfc0007fc, 0xf00002ac, 0x1f0000, // VSX Vector Round to Single-Precision Integer Exact using Current rounding mode XX2-form (xvrspic XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVRSPIM, 0xfc0007fc, 0xf00002e4, 0x1f0000, // VSX Vector Round to Single-Precision Integer using round toward -Infinity XX2-form (xvrspim XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVRSPIP, 0xfc0007fc, 0xf00002a4, 0x1f0000, // VSX Vector Round to Single-Precision Integer using round toward +Infinity XX2-form (xvrspip XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVRSPIZ, 0xfc0007fc, 0xf0000264, 0x1f0000, // VSX Vector Round to Single-Precision Integer using round toward Zero XX2-form (xvrspiz XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVRSQRTEDP, 0xfc0007fc, 0xf0000328, 0x1f0000, // VSX Vector Reciprocal Square Root Estimate Double-Precision XX2-form (xvrsqrtedp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVRSQRTESP, 0xfc0007fc, 0xf0000228, 0x1f0000, // VSX Vector Reciprocal Square Root Estimate Single-Precision XX2-form (xvrsqrtesp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVSQRTDP, 0xfc0007fc, 0xf000032c, 0x1f0000, // VSX Vector Square Root Double-Precision XX2-form (xvsqrtdp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVSQRTSP, 0xfc0007fc, 0xf000022c, 0x1f0000, // VSX Vector Square Root Single-Precision XX2-form (xvsqrtsp XT,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20}},
	{XVSUBDP, 0xfc0007f8, 0xf0000340, 0x0, // VSX Vector Subtract Double-Precision XX3-form (xvsubdp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XVSUBSP, 0xfc0007f8, 0xf0000240, 0x0, // VSX Vector Subtract Single-Precision XX3-form (xvsubsp XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XVTDIVDP, 0xfc0007f8, 0xf00003e8, 0x600001, // VSX Vector Test for software Divide Double-Precision XX3-form (xvtdivdp BF,XA,XB)
		[5]*argField{ap_CondRegField_6_8, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XVTDIVSP, 0xfc0007f8, 0xf00002e8, 0x600001, // VSX Vector Test for software Divide Single-Precision XX3-form (xvtdivsp BF,XA,XB)
		[5]*argField{ap_CondRegField_6_8, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XVTSQRTDP, 0xfc0007fc, 0xf00003a8, 0x7f0001, // VSX Vector Test for software Square Root Double-Precision XX2-form (xvtsqrtdp BF,XB)
		[5]*argField{ap_CondRegField_6_8, ap_VecReg_30_30_16_20}},
	{XVTSQRTSP, 0xfc0007fc, 0xf00002a8, 0x7f0001, // VSX Vector Test for software Square Root Single-Precision XX2-form (xvtsqrtsp BF,XB)
		[5]*argField{ap_CondRegField_6_8, ap_VecReg_30_30_16_20}},
	{XXLAND, 0xfc0007f8, 0xf0000410, 0x0, // VSX Logical AND XX3-form (xxland XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XXLANDC, 0xfc0007f8, 0xf0000450, 0x0, // VSX Logical AND with Complement XX3-form (xxlandc XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XXLEQV, 0xfc0007f8, 0xf00005d0, 0x0, // VSX Logical Equivalence XX3-form (xxleqv XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XXLNAND, 0xfc0007f8, 0xf0000590, 0x0, // VSX Logical NAND XX3-form (xxlnand XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XXLORC, 0xfc0007f8, 0xf0000550, 0x0, // VSX Logical OR with Complement XX3-form (xxlorc XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XXLNOR, 0xfc0007f8, 0xf0000510, 0x0, // VSX Logical NOR XX3-form (xxlnor XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XXLOR, 0xfc0007f8, 0xf0000490, 0x0, // VSX Logical OR XX3-form (xxlor XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XXLXOR, 0xfc0007f8, 0xf00004d0, 0x0, // VSX Logical XOR XX3-form (xxlxor XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XXMRGHW, 0xfc0007f8, 0xf0000090, 0x0, // VSX Merge High Word XX3-form (xxmrghw XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XXMRGLW, 0xfc0007f8, 0xf0000190, 0x0, // VSX Merge Low Word XX3-form (xxmrglw XT,XA,XB)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20}},
	{XXPERMDI, 0xfc0004f8, 0xf0000050, 0x0, // VSX Permute Doubleword Immediate XX3-form (xxpermdi XT,XA,XB,DM)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20, ap_ImmUnsigned_22_23}},
	{XXSEL, 0xfc000030, 0xf0000030, 0x0, // VSX Select XX4-form (xxsel XT,XA,XB,XC)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20, ap_VecReg_28_28_21_25}},
	{XXSLDWI, 0xfc0004f8, 0xf0000010, 0x0, // VSX Shift Left Double by Word Immediate XX3-form (xxsldwi XT,XA,XB,SHW)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_29_29_11_15, ap_VecReg_30_30_16_20, ap_ImmUnsigned_22_23}},
	{XXSPLTW, 0xfc0007fc, 0xf0000290, 0x1c0000, // VSX Splat Word XX2-form (xxspltw XT,XB,UIM)
		[5]*argField{ap_VecReg_31_31_6_10, ap_VecReg_30_30_16_20, ap_ImmUnsigned_14_15}},
	{BRINC, 0xfc0007ff, 0x1000020f, 0x0, // Bit Reversed Increment EVX-form (brinc RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVABS, 0xfc0007ff, 0x10000208, 0xf800, // Vector Absolute Value EVX-form (evabs RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVADDIW, 0xfc0007ff, 0x10000202, 0x0, // Vector Add Immediate Word EVX-form (evaddiw RT,RB,UI)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20, ap_ImmUnsigned_11_15}},
	{EVADDSMIAAW, 0xfc0007ff, 0x100004c9, 0xf800, // Vector Add Signed, Modulo, Integer to Accumulator Word EVX-form (evaddsmiaaw RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVADDSSIAAW, 0xfc0007ff, 0x100004c1, 0xf800, // Vector Add Signed, Saturate, Integer to Accumulator Word EVX-form (evaddssiaaw RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVADDUMIAAW, 0xfc0007ff, 0x100004c8, 0xf800, // Vector Add Unsigned, Modulo, Integer to Accumulator Word EVX-form (evaddumiaaw RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVADDUSIAAW, 0xfc0007ff, 0x100004c0, 0xf800, // Vector Add Unsigned, Saturate, Integer to Accumulator Word EVX-form (evaddusiaaw RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVADDW, 0xfc0007ff, 0x10000200, 0x0, // Vector Add Word EVX-form (evaddw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVAND, 0xfc0007ff, 0x10000211, 0x0, // Vector AND EVX-form (evand RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVCMPEQ, 0xfc0007ff, 0x10000234, 0x600000, // Vector Compare Equal EVX-form (evcmpeq BF,RA,RB)
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EVANDC, 0xfc0007ff, 0x10000212, 0x0, // Vector AND with Complement EVX-form (evandc RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVCMPGTS, 0xfc0007ff, 0x10000231, 0x600000, // Vector Compare Greater Than Signed EVX-form (evcmpgts BF,RA,RB)
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EVCMPGTU, 0xfc0007ff, 0x10000230, 0x600000, // Vector Compare Greater Than Unsigned EVX-form (evcmpgtu BF,RA,RB)
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EVCMPLTU, 0xfc0007ff, 0x10000232, 0x600000, // Vector Compare Less Than Unsigned EVX-form (evcmpltu BF,RA,RB)
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EVCMPLTS, 0xfc0007ff, 0x10000233, 0x600000, // Vector Compare Less Than Signed EVX-form (evcmplts BF,RA,RB)
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EVCNTLSW, 0xfc0007ff, 0x1000020e, 0xf800, // Vector Count Leading Signed Bits Word EVX-form (evcntlsw RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVCNTLZW, 0xfc0007ff, 0x1000020d, 0xf800, // Vector Count Leading Zeros Word EVX-form (evcntlzw RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVDIVWS, 0xfc0007ff, 0x100004c6, 0x0, // Vector Divide Word Signed EVX-form (evdivws RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVDIVWU, 0xfc0007ff, 0x100004c7, 0x0, // Vector Divide Word Unsigned EVX-form (evdivwu RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVEQV, 0xfc0007ff, 0x10000219, 0x0, // Vector Equivalent EVX-form (eveqv RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVEXTSB, 0xfc0007ff, 0x1000020a, 0xf800, // Vector Extend Sign Byte EVX-form (evextsb RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVEXTSH, 0xfc0007ff, 0x1000020b, 0xf800, // Vector Extend Sign Halfword EVX-form (evextsh RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVLDD, 0xfc0007ff, 0x10000301, 0x0, // Vector Load Double Word into Double Word EVX-form (evldd RT,D(RA))
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_Reg_11_15}},
	{EVLDH, 0xfc0007ff, 0x10000305, 0x0, // Vector Load Double into Four Halfwords EVX-form (evldh RT,D(RA))
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_Reg_11_15}},
	{EVLDDX, 0xfc0007ff, 0x10000300, 0x0, // Vector Load Double Word into Double Word Indexed EVX-form (evlddx RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVLDHX, 0xfc0007ff, 0x10000304, 0x0, // Vector Load Double into Four Halfwords Indexed EVX-form (evldhx RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVLDW, 0xfc0007ff, 0x10000303, 0x0, // Vector Load Double into Two Words EVX-form (evldw RT,D(RA))
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_Reg_11_15}},
	{EVLHHESPLAT, 0xfc0007ff, 0x10000309, 0x0, // Vector Load Halfword into Halfwords Even and Splat EVX-form (evlhhesplat RT,D(RA))
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_Reg_11_15}},
	{EVLDWX, 0xfc0007ff, 0x10000302, 0x0, // Vector Load Double into Two Words Indexed EVX-form (evldwx RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVLHHESPLATX, 0xfc0007ff, 0x10000308, 0x0, // Vector Load Halfword into Halfwords Even and Splat Indexed EVX-form (evlhhesplatx RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVLHHOSSPLAT, 0xfc0007ff, 0x1000030f, 0x0, // Vector Load Halfword into Halfword Odd Signed and Splat EVX-form (evlhhossplat RT,D(RA))
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_Reg_11_15}},
	{EVLHHOUSPLAT, 0xfc0007ff, 0x1000030d, 0x0, // Vector Load Halfword into Halfword Odd Unsigned and Splat EVX-form (evlhhousplat RT,D(RA))
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_Reg_11_15}},
	{EVLHHOSSPLATX, 0xfc0007ff, 0x1000030e, 0x0, // Vector Load Halfword into Halfword Odd Signed and Splat Indexed EVX-form (evlhhossplatx RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVLHHOUSPLATX, 0xfc0007ff, 0x1000030c, 0x0, // Vector Load Halfword into Halfword Odd Unsigned and Splat Indexed EVX-form (evlhhousplatx RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVLWHE, 0xfc0007ff, 0x10000311, 0x0, // Vector Load Word into Two Halfwords Even EVX-form (evlwhe RT,D(RA))
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_Reg_11_15}},
	{EVLWHOS, 0xfc0007ff, 0x10000317, 0x0, // Vector Load Word into Two Halfwords Odd Signed (with sign extension) EVX-form (evlwhos RT,D(RA))
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_Reg_11_15}},
	{EVLWHEX, 0xfc0007ff, 0x10000310, 0x0, // Vector Load Word into Two Halfwords Even Indexed EVX-form (evlwhex RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVLWHOSX, 0xfc0007ff, 0x10000316, 0x0, // Vector Load Word into Two Halfwords Odd Signed Indexed (with sign extension) EVX-form (evlwhosx RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVLWHOU, 0xfc0007ff, 0x10000315, 0x0, // Vector Load Word into Two Halfwords Odd Unsigned (zero-extended) EVX-form (evlwhou RT,D(RA))
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_Reg_11_15}},
	{EVLWHSPLAT, 0xfc0007ff, 0x1000031d, 0x0, // Vector Load Word into Two Halfwords and Splat EVX-form (evlwhsplat RT,D(RA))
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_Reg_11_15}},
	{EVLWHOUX, 0xfc0007ff, 0x10000314, 0x0, // Vector Load Word into Two Halfwords Odd Unsigned Indexed (zero-extended) EVX-form (evlwhoux RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVLWHSPLATX, 0xfc0007ff, 0x1000031c, 0x0, // Vector Load Word into Two Halfwords and Splat Indexed EVX-form (evlwhsplatx RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVLWWSPLAT, 0xfc0007ff, 0x10000319, 0x0, // Vector Load Word into Word and Splat EVX-form (evlwwsplat RT,D(RA))
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_Reg_11_15}},
	{EVMERGEHI, 0xfc0007ff, 0x1000022c, 0x0, // Vector Merge High EVX-form (evmergehi RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVLWWSPLATX, 0xfc0007ff, 0x10000318, 0x0, // Vector Load Word into Word and Splat Indexed EVX-form (evlwwsplatx RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMERGELO, 0xfc0007ff, 0x1000022d, 0x0, // Vector Merge Low EVX-form (evmergelo RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMERGEHILO, 0xfc0007ff, 0x1000022e, 0x0, // Vector Merge High/Low EVX-form (evmergehilo RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHEGSMFAA, 0xfc0007ff, 0x1000052b, 0x0, // Vector Multiply Halfwords, Even, Guarded, Signed, Modulo, Fractional and Accumulate EVX-form (evmhegsmfaa RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMERGELOHI, 0xfc0007ff, 0x1000022f, 0x0, // Vector Merge Low/High EVX-form (evmergelohi RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHEGSMFAN, 0xfc0007ff, 0x100005ab, 0x0, // Vector Multiply Halfwords, Even, Guarded, Signed, Modulo, Fractional and Accumulate Negative EVX-form (evmhegsmfan RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHEGSMIAA, 0xfc0007ff, 0x10000529, 0x0, // Vector Multiply Halfwords, Even, Guarded, Signed, Modulo, Integer and Accumulate EVX-form (evmhegsmiaa RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHEGUMIAA, 0xfc0007ff, 0x10000528, 0x0, // Vector Multiply Halfwords, Even, Guarded, Unsigned, Modulo, Integer and Accumulate EVX-form (evmhegumiaa RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHEGSMIAN, 0xfc0007ff, 0x100005a9, 0x0, // Vector Multiply Halfwords, Even, Guarded, Signed, Modulo, Integer and Accumulate Negative EVX-form (evmhegsmian RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHEGUMIAN, 0xfc0007ff, 0x100005a8, 0x0, // Vector Multiply Halfwords, Even, Guarded, Unsigned, Modulo, Integer and Accumulate Negative EVX-form (evmhegumian RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHESMF, 0xfc0007ff, 0x1000040b, 0x0, // Vector Multiply Halfwords, Even, Signed, Modulo, Fractional EVX-form (evmhesmf RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHESMFAAW, 0xfc0007ff, 0x1000050b, 0x0, // Vector Multiply Halfwords, Even, Signed, Modulo, Fractional and Accumulate into Words EVX-form (evmhesmfaaw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHESMFA, 0xfc0007ff, 0x1000042b, 0x0, // Vector Multiply Halfwords, Even, Signed, Modulo, Fractional to Accumulator EVX-form (evmhesmfa RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHESMFANW, 0xfc0007ff, 0x1000058b, 0x0, // Vector Multiply Halfwords, Even, Signed, Modulo, Fractional and Accumulate Negative into Words EVX-form (evmhesmfanw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHESMI, 0xfc0007ff, 0x10000409, 0x0, // Vector Multiply Halfwords, Even, Signed, Modulo, Integer EVX-form (evmhesmi RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHESMIAAW, 0xfc0007ff, 0x10000509, 0x0, // Vector Multiply Halfwords, Even, Signed, Modulo, Integer and Accumulate into Words EVX-form (evmhesmiaaw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHESMIA, 0xfc0007ff, 0x10000429, 0x0, // Vector Multiply Halfwords, Even, Signed, Modulo, Integer to Accumulator EVX-form (evmhesmia RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHESMIANW, 0xfc0007ff, 0x10000589, 0x0, // Vector Multiply Halfwords, Even, Signed, Modulo, Integer and Accumulate Negative into Words EVX-form (evmhesmianw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHESSF, 0xfc0007ff, 0x10000403, 0x0, // Vector Multiply Halfwords, Even, Signed, Saturate, Fractional EVX-form (evmhessf RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHESSFA, 0xfc0007ff, 0x10000423, 0x0, // Vector Multiply Halfwords, Even, Signed, Saturate, Fractional to Accumulator EVX-form (evmhessfa RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHESSFAAW, 0xfc0007ff, 0x10000503, 0x0, // Vector Multiply Halfwords, Even, Signed, Saturate, Fractional and Accumulate into Words EVX-form (evmhessfaaw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHESSFANW, 0xfc0007ff, 0x10000583, 0x0, // Vector Multiply Halfwords, Even, Signed, Saturate, Fractional and Accumulate Negative into Words EVX-form (evmhessfanw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHESSIAAW, 0xfc0007ff, 0x10000501, 0x0, // Vector Multiply Halfwords, Even, Signed, Saturate, Integer and Accumulate into Words EVX-form (evmhessiaaw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHESSIANW, 0xfc0007ff, 0x10000581, 0x0, // Vector Multiply Halfwords, Even, Signed, Saturate, Integer and Accumulate Negative into Words EVX-form (evmhessianw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHEUMI, 0xfc0007ff, 0x10000408, 0x0, // Vector Multiply Halfwords, Even, Unsigned, Modulo, Integer EVX-form (evmheumi RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHEUMIAAW, 0xfc0007ff, 0x10000508, 0x0, // Vector Multiply Halfwords, Even, Unsigned, Modulo, Integer and Accumulate into Words EVX-form (evmheumiaaw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHEUMIA, 0xfc0007ff, 0x10000428, 0x0, // Vector Multiply Halfwords, Even, Unsigned, Modulo, Integer to Accumulator EVX-form (evmheumia RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHEUMIANW, 0xfc0007ff, 0x10000588, 0x0, // Vector Multiply Halfwords, Even, Unsigned, Modulo, Integer and Accumulate Negative into Words EVX-form (evmheumianw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHEUSIAAW, 0xfc0007ff, 0x10000500, 0x0, // Vector Multiply Halfwords, Even, Unsigned, Saturate, Integer and Accumulate into Words EVX-form (evmheusiaaw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHEUSIANW, 0xfc0007ff, 0x10000580, 0x0, // Vector Multiply Halfwords, Even, Unsigned, Saturate, Integer and Accumulate Negative into Words EVX-form (evmheusianw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOGSMFAA, 0xfc0007ff, 0x1000052f, 0x0, // Vector Multiply Halfwords, Odd, Guarded, Signed, Modulo, Fractional and Accumulate EVX-form (evmhogsmfaa RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOGSMIAA, 0xfc0007ff, 0x1000052d, 0x0, // Vector Multiply Halfwords, Odd, Guarded, Signed, Modulo, Integer and Accumulate EVX-form (evmhogsmiaa RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOGSMFAN, 0xfc0007ff, 0x100005af, 0x0, // Vector Multiply Halfwords, Odd, Guarded, Signed, Modulo, Fractional and Accumulate Negative EVX-form (evmhogsmfan RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOGSMIAN, 0xfc0007ff, 0x100005ad, 0x0, // Vector Multiply Halfwords, Odd, Guarded, Signed, Modulo, Integer and Accumulate Negative EVX-form (evmhogsmian RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOGUMIAA, 0xfc0007ff, 0x1000052c, 0x0, // Vector Multiply Halfwords, Odd, Guarded, Unsigned, Modulo, Integer and Accumulate EVX-form (evmhogumiaa RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOSMF, 0xfc0007ff, 0x1000040f, 0x0, // Vector Multiply Halfwords, Odd, Signed, Modulo, Fractional EVX-form (evmhosmf RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOGUMIAN, 0xfc0007ff, 0x100005ac, 0x0, // Vector Multiply Halfwords, Odd, Guarded, Unsigned, Modulo, Integer and Accumulate Negative EVX-form (evmhogumian RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOSMFA, 0xfc0007ff, 0x1000042f, 0x0, // Vector Multiply Halfwords, Odd, Signed, Modulo, Fractional to Accumulator EVX-form (evmhosmfa RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOSMFAAW, 0xfc0007ff, 0x1000050f, 0x0, // Vector Multiply Halfwords, Odd, Signed, Modulo, Fractional and Accumulate into Words EVX-form (evmhosmfaaw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOSMI, 0xfc0007ff, 0x1000040d, 0x0, // Vector Multiply Halfwords, Odd, Signed, Modulo, Integer EVX-form (evmhosmi RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOSMFANW, 0xfc0007ff, 0x1000058f, 0x0, // Vector Multiply Halfwords, Odd, Signed, Modulo, Fractional and Accumulate Negative into Words EVX-form (evmhosmfanw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOSMIA, 0xfc0007ff, 0x1000042d, 0x0, // Vector Multiply Halfwords, Odd, Signed, Modulo, Integer to Accumulator EVX-form (evmhosmia RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOSMIAAW, 0xfc0007ff, 0x1000050d, 0x0, // Vector Multiply Halfwords, Odd, Signed, Modulo, Integer and Accumulate into Words EVX-form (evmhosmiaaw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOSMIANW, 0xfc0007ff, 0x1000058d, 0x0, // Vector Multiply Halfwords, Odd, Signed, Modulo, Integer and Accumulate Negative into Words EVX-form (evmhosmianw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOSSF, 0xfc0007ff, 0x10000407, 0x0, // Vector Multiply Halfwords, Odd, Signed, Saturate, Fractional EVX-form (evmhossf RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOSSFA, 0xfc0007ff, 0x10000427, 0x0, // Vector Multiply Halfwords, Odd, Signed, Saturate, Fractional to Accumulator EVX-form (evmhossfa RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOSSFAAW, 0xfc0007ff, 0x10000507, 0x0, // Vector Multiply Halfwords, Odd, Signed, Saturate, Fractional and Accumulate into Words EVX-form (evmhossfaaw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOSSFANW, 0xfc0007ff, 0x10000587, 0x0, // Vector Multiply Halfwords, Odd, Signed, Saturate, Fractional and Accumulate Negative into Words EVX-form (evmhossfanw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOSSIAAW, 0xfc0007ff, 0x10000505, 0x0, // Vector Multiply Halfwords, Odd, Signed, Saturate, Integer and Accumulate into Words EVX-form (evmhossiaaw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOUMI, 0xfc0007ff, 0x1000040c, 0x0, // Vector Multiply Halfwords, Odd, Unsigned, Modulo, Integer EVX-form (evmhoumi RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOSSIANW, 0xfc0007ff, 0x10000585, 0x0, // Vector Multiply Halfwords, Odd, Signed, Saturate, Integer and Accumulate Negative into Words EVX-form (evmhossianw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOUMIA, 0xfc0007ff, 0x1000042c, 0x0, // Vector Multiply Halfwords, Odd, Unsigned, Modulo, Integer to Accumulator EVX-form (evmhoumia RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOUMIAAW, 0xfc0007ff, 0x1000050c, 0x0, // Vector Multiply Halfwords, Odd, Unsigned, Modulo, Integer and Accumulate into Words EVX-form (evmhoumiaaw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOUSIAAW, 0xfc0007ff, 0x10000504, 0x0, // Vector Multiply Halfwords, Odd, Unsigned, Saturate, Integer and Accumulate into Words EVX-form (evmhousiaaw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOUMIANW, 0xfc0007ff, 0x1000058c, 0x0, // Vector Multiply Halfwords, Odd, Unsigned, Modulo, Integer and Accumulate Negative into Words EVX-form (evmhoumianw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOUSIANW, 0xfc0007ff, 0x10000584, 0x0, // Vector Multiply Halfwords, Odd, Unsigned, Saturate, Integer and Accumulate Negative into Words EVX-form (evmhousianw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMRA, 0xfc0007ff, 0x100004c4, 0xf800, // Initialize Accumulator EVX-form (evmra RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVMWHSMF, 0xfc0007ff, 0x1000044f, 0x0, // Vector Multiply Word High Signed, Modulo, Fractional EVX-form (evmwhsmf RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWHSMI, 0xfc0007ff, 0x1000044d, 0x0, // Vector Multiply Word High Signed, Modulo, Integer EVX-form (evmwhsmi RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWHSMFA, 0xfc0007ff, 0x1000046f, 0x0, // Vector Multiply Word High Signed, Modulo, Fractional to Accumulator EVX-form (evmwhsmfa RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWHSMIA, 0xfc0007ff, 0x1000046d, 0x0, // Vector Multiply Word High Signed, Modulo, Integer to Accumulator EVX-form (evmwhsmia RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWHSSF, 0xfc0007ff, 0x10000447, 0x0, // Vector Multiply Word High Signed, Saturate, Fractional EVX-form (evmwhssf RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWHUMI, 0xfc0007ff, 0x1000044c, 0x0, // Vector Multiply Word High Unsigned, Modulo, Integer EVX-form (evmwhumi RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWHSSFA, 0xfc0007ff, 0x10000467, 0x0, // Vector Multiply Word High Signed, Saturate, Fractional to Accumulator EVX-form (evmwhssfa RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWHUMIA, 0xfc0007ff, 0x1000046c, 0x0, // Vector Multiply Word High Unsigned, Modulo, Integer to Accumulator EVX-form (evmwhumia RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWLSMIAAW, 0xfc0007ff, 0x10000549, 0x0, // Vector Multiply Word Low Signed, Modulo, Integer and Accumulate into Words EVX-form (evmwlsmiaaw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWLSSIAAW, 0xfc0007ff, 0x10000541, 0x0, // Vector Multiply Word Low Signed, Saturate, Integer and Accumulate into Words EVX-form (evmwlssiaaw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWLSMIANW, 0xfc0007ff, 0x100005c9, 0x0, // Vector Multiply Word Low Signed, Modulo, Integer and Accumulate Negative in Words EVX-form (evmwlsmianw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWLSSIANW, 0xfc0007ff, 0x100005c1, 0x0, // Vector Multiply Word Low Signed, Saturate, Integer and Accumulate Negative in Words EVX-form (evmwlssianw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWLUMI, 0xfc0007ff, 0x10000448, 0x0, // Vector Multiply Word Low Unsigned, Modulo, Integer EVX-form (evmwlumi RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWLUMIAAW, 0xfc0007ff, 0x10000548, 0x0, // Vector Multiply Word Low Unsigned, Modulo, Integer and Accumulate into Words EVX-form (evmwlumiaaw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWLUMIA, 0xfc0007ff, 0x10000468, 0x0, // Vector Multiply Word Low Unsigned, Modulo, Integer to Accumulator EVX-form (evmwlumia RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWLUMIANW, 0xfc0007ff, 0x100005c8, 0x0, // Vector Multiply Word Low Unsigned, Modulo, Integer and Accumulate Negative in Words EVX-form (evmwlumianw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWLUSIAAW, 0xfc0007ff, 0x10000540, 0x0, // Vector Multiply Word Low Unsigned, Saturate, Integer and Accumulate into Words EVX-form (evmwlusiaaw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWSMF, 0xfc0007ff, 0x1000045b, 0x0, // Vector Multiply Word Signed, Modulo, Fractional EVX-form (evmwsmf RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWLUSIANW, 0xfc0007ff, 0x100005c0, 0x0, // Vector Multiply Word Low Unsigned, Saturate, Integer and Accumulate Negative in Words EVX-form (evmwlusianw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWSMFA, 0xfc0007ff, 0x1000047b, 0x0, // Vector Multiply Word Signed, Modulo, Fractional to Accumulator EVX-form (evmwsmfa RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWSMFAA, 0xfc0007ff, 0x1000055b, 0x0, // Vector Multiply Word Signed, Modulo, Fractional and Accumulate EVX-form (evmwsmfaa RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWSMI, 0xfc0007ff, 0x10000459, 0x0, // Vector Multiply Word Signed, Modulo, Integer EVX-form (evmwsmi RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWSMIAA, 0xfc0007ff, 0x10000559, 0x0, // Vector Multiply Word Signed, Modulo, Integer and Accumulate EVX-form (evmwsmiaa RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWSMFAN, 0xfc0007ff, 0x100005db, 0x0, // Vector Multiply Word Signed, Modulo, Fractional and Accumulate Negative EVX-form (evmwsmfan RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWSMIA, 0xfc0007ff, 0x10000479, 0x0, // Vector Multiply Word Signed, Modulo, Integer to Accumulator EVX-form (evmwsmia RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWSMIAN, 0xfc0007ff, 0x100005d9, 0x0, // Vector Multiply Word Signed, Modulo, Integer and Accumulate Negative EVX-form (evmwsmian RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWSSF, 0xfc0007ff, 0x10000453, 0x0, // Vector Multiply Word Signed, Saturate, Fractional EVX-form (evmwssf RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWSSFA, 0xfc0007ff, 0x10000473, 0x0, // Vector Multiply Word Signed, Saturate, Fractional to Accumulator EVX-form (evmwssfa RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWSSFAA, 0xfc0007ff, 0x10000553, 0x0, // Vector Multiply Word Signed, Saturate, Fractional and Accumulate EVX-form (evmwssfaa RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWUMI, 0xfc0007ff, 0x10000458, 0x0, // Vector Multiply Word Unsigned, Modulo, Integer EVX-form (evmwumi RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWSSFAN, 0xfc0007ff, 0x100005d3, 0x0, // Vector Multiply Word Signed, Saturate, Fractional and Accumulate Negative EVX-form (evmwssfan RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWUMIA, 0xfc0007ff, 0x10000478, 0x0, // Vector Multiply Word Unsigned, Modulo, Integer to Accumulator EVX-form (evmwumia RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWUMIAA, 0xfc0007ff, 0x10000558, 0x0, // Vector Multiply Word Unsigned, Modulo, Integer and Accumulate EVX-form (evmwumiaa RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVNAND, 0xfc0007ff, 0x1000021e, 0x0, // Vector NAND EVX-form (evnand RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWUMIAN, 0xfc0007ff, 0x100005d8, 0x0, // Vector Multiply Word Unsigned, Modulo, Integer and Accumulate Negative EVX-form (evmwumian RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVNEG, 0xfc0007ff, 0x10000209, 0xf800, // Vector Negate EVX-form (evneg RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVNOR, 0xfc0007ff, 0x10000218, 0x0, // Vector NOR EVX-form (evnor RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVORC, 0xfc0007ff, 0x1000021b, 0x0, // Vector OR with Complement EVX-form (evorc RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVOR, 0xfc0007ff, 0x10000217, 0x0, // Vector OR EVX-form (evor RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVRLW, 0xfc0007ff, 0x10000228, 0x0, // Vector Rotate Left Word EVX-form (evrlw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVRLWI, 0xfc0007ff, 0x1000022a, 0x0, // Vector Rotate Left Word Immediate EVX-form (evrlwi RT,RA,UI)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_ImmUnsigned_16_20}},
	{EVSEL, 0xfc0007f8, 0x10000278, 0x0, // Vector Select EVS-form (evsel RT,RA,RB,BFA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20, ap_CondRegField_29_31}},
	{EVRNDW, 0xfc0007ff, 0x1000020c, 0xf800, // Vector Round Word EVX-form (evrndw RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVSLW, 0xfc0007ff, 0x10000224, 0x0, // Vector Shift Left Word EVX-form (evslw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVSPLATFI, 0xfc0007ff, 0x1000022b, 0xf800, // Vector Splat Fractional Immediate EVX-form (evsplatfi RT,SI)
		[5]*argField{ap_Reg_6_10, ap_ImmSigned_11_15}},
	{EVSRWIS, 0xfc0007ff, 0x10000223, 0x0, // Vector Shift Right Word Immediate Signed EVX-form (evsrwis RT,RA,UI)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_ImmUnsigned_16_20}},
	{EVSLWI, 0xfc0007ff, 0x10000226, 0x0, // Vector Shift Left Word Immediate EVX-form (evslwi RT,RA,UI)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_ImmUnsigned_16_20}},
	{EVSPLATI, 0xfc0007ff, 0x10000229, 0xf800, // Vector Splat Immediate EVX-form (evsplati RT,SI)
		[5]*argField{ap_Reg_6_10, ap_ImmSigned_11_15}},
	{EVSRWIU, 0xfc0007ff, 0x10000222, 0x0, // Vector Shift Right Word Immediate Unsigned EVX-form (evsrwiu RT,RA,UI)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_ImmUnsigned_16_20}},
	{EVSRWS, 0xfc0007ff, 0x10000221, 0x0, // Vector Shift Right Word Signed EVX-form (evsrws RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVSTDD, 0xfc0007ff, 0x10000321, 0x0, // Vector Store Double of Double EVX-form (evstdd RS,D(RA))
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_Reg_11_15}},
	{EVSRWU, 0xfc0007ff, 0x10000220, 0x0, // Vector Shift Right Word Unsigned EVX-form (evsrwu RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVSTDDX, 0xfc0007ff, 0x10000320, 0x0, // Vector Store Double of Double Indexed EVX-form (evstddx RS,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVSTDH, 0xfc0007ff, 0x10000325, 0x0, // Vector Store Double of Four Halfwords EVX-form (evstdh RS,D(RA))
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_Reg_11_15}},
	{EVSTDW, 0xfc0007ff, 0x10000323, 0x0, // Vector Store Double of Two Words EVX-form (evstdw RS,D(RA))
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_Reg_11_15}},
	{EVSTDHX, 0xfc0007ff, 0x10000324, 0x0, // Vector Store Double of Four Halfwords Indexed EVX-form (evstdhx RS,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVSTDWX, 0xfc0007ff, 0x10000322, 0x0, // Vector Store Double of Two Words Indexed EVX-form (evstdwx RS,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVSTWHE, 0xfc0007ff, 0x10000331, 0x0, // Vector Store Word of Two Halfwords from Even EVX-form (evstwhe RS,D(RA))
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_Reg_11_15}},
	{EVSTWHO, 0xfc0007ff, 0x10000335, 0x0, // Vector Store Word of Two Halfwords from Odd EVX-form (evstwho RS,D(RA))
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_Reg_11_15}},
	{EVSTWWE, 0xfc0007ff, 0x10000339, 0x0, // Vector Store Word of Word from Even EVX-form (evstwwe RS,D(RA))
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_Reg_11_15}},
	{EVSTWHEX, 0xfc0007ff, 0x10000330, 0x0, // Vector Store Word of Two Halfwords from Even Indexed EVX-form (evstwhex RS,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVSTWHOX, 0xfc0007ff, 0x10000334, 0x0, // Vector Store Word of Two Halfwords from Odd Indexed EVX-form (evstwhox RS,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVSTWWEX, 0xfc0007ff, 0x10000338, 0x0, // Vector Store Word of Word from Even Indexed EVX-form (evstwwex RS,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVSTWWO, 0xfc0007ff, 0x1000033d, 0x0, // Vector Store Word of Word from Odd EVX-form (evstwwo RS,D(RA))
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_Reg_11_15}},
	{EVSUBFSMIAAW, 0xfc0007ff, 0x100004cb, 0xf800, // Vector Subtract Signed, Modulo, Integer to Accumulator Word EVX-form (evsubfsmiaaw RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVSTWWOX, 0xfc0007ff, 0x1000033c, 0x0, // Vector Store Word of Word from Odd Indexed EVX-form (evstwwox RS,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVSUBFSSIAAW, 0xfc0007ff, 0x100004c3, 0xf800, // Vector Subtract Signed, Saturate, Integer to Accumulator Word EVX-form (evsubfssiaaw RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVSUBFUMIAAW, 0xfc0007ff, 0x100004ca, 0xf800, // Vector Subtract Unsigned, Modulo, Integer to Accumulator Word EVX-form (evsubfumiaaw RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVSUBFUSIAAW, 0xfc0007ff, 0x100004c2, 0xf800, // Vector Subtract Unsigned, Saturate, Integer to Accumulator Word EVX-form (evsubfusiaaw RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVSUBFW, 0xfc0007ff, 0x10000204, 0x0, // Vector Subtract from Word EVX-form (evsubfw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVSUBIFW, 0xfc0007ff, 0x10000206, 0x0, // Vector Subtract Immediate from Word EVX-form (evsubifw RT,UI,RB)
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_11_15, ap_Reg_16_20}},
	{EVXOR, 0xfc0007ff, 0x10000216, 0x0, // Vector XOR EVX-form (evxor RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVFSABS, 0xfc0007ff, 0x10000284, 0xf800, // Vector Floating-Point Single-Precision Absolute Value EVX-form (evfsabs RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVFSNABS, 0xfc0007ff, 0x10000285, 0xf800, // Vector Floating-Point Single-Precision Negative Absolute Value EVX-form (evfsnabs RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVFSNEG, 0xfc0007ff, 0x10000286, 0xf800, // Vector Floating-Point Single-Precision Negate EVX-form (evfsneg RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVFSADD, 0xfc0007ff, 0x10000280, 0x0, // Vector Floating-Point Single-Precision Add EVX-form (evfsadd RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVFSMUL, 0xfc0007ff, 0x10000288, 0x0, // Vector Floating-Point Single-Precision Multiply EVX-form (evfsmul RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVFSSUB, 0xfc0007ff, 0x10000281, 0x0, // Vector Floating-Point Single-Precision Subtract EVX-form (evfssub RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVFSDIV, 0xfc0007ff, 0x10000289, 0x0, // Vector Floating-Point Single-Precision Divide EVX-form (evfsdiv RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVFSCMPGT, 0xfc0007ff, 0x1000028c, 0x600000, // Vector Floating-Point Single-Precision Compare Greater Than EVX-form (evfscmpgt BF,RA,RB)
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EVFSCMPLT, 0xfc0007ff, 0x1000028d, 0x600000, // Vector Floating-Point Single-Precision Compare Less Than EVX-form (evfscmplt BF,RA,RB)
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EVFSCMPEQ, 0xfc0007ff, 0x1000028e, 0x600000, // Vector Floating-Point Single-Precision Compare Equal EVX-form (evfscmpeq BF,RA,RB)
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EVFSTSTGT, 0xfc0007ff, 0x1000029c, 0x600000, // Vector Floating-Point Single-Precision Test Greater Than EVX-form (evfststgt BF,RA,RB)
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EVFSTSTLT, 0xfc0007ff, 0x1000029d, 0x600000, // Vector Floating-Point Single-Precision Test Less Than EVX-form (evfststlt BF,RA,RB)
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EVFSTSTEQ, 0xfc0007ff, 0x1000029e, 0x600000, // Vector Floating-Point Single-Precision Test Equal EVX-form (evfststeq BF,RA,RB)
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EVFSCFSI, 0xfc0007ff, 0x10000291, 0x1f0000, // Vector Convert Floating-Point Single-Precision from Signed Integer EVX-form (evfscfsi RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EVFSCFSF, 0xfc0007ff, 0x10000293, 0x1f0000, // Vector Convert Floating-Point Single-Precision from Signed Fraction EVX-form (evfscfsf RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EVFSCFUI, 0xfc0007ff, 0x10000290, 0x1f0000, // Vector Convert Floating-Point Single-Precision from Unsigned Integer EVX-form (evfscfui RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EVFSCFUF, 0xfc0007ff, 0x10000292, 0x1f0000, // Vector Convert Floating-Point Single-Precision from Unsigned Fraction EVX-form (evfscfuf RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EVFSCTSI, 0xfc0007ff, 0x10000295, 0x1f0000, // Vector Convert Floating-Point Single-Precision to Signed Integer EVX-form (evfsctsi RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EVFSCTUI, 0xfc0007ff, 0x10000294, 0x1f0000, // Vector Convert Floating-Point Single-Precision to Unsigned Integer EVX-form (evfsctui RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EVFSCTSIZ, 0xfc0007ff, 0x1000029a, 0x1f0000, // Vector Convert Floating-Point Single-Precision to Signed Integer with Round toward Zero EVX-form (evfsctsiz RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EVFSCTUIZ, 0xfc0007ff, 0x10000298, 0x1f0000, // Vector Convert Floating-Point Single-Precision to Unsigned Integer with Round toward Zero EVX-form (evfsctuiz RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EVFSCTSF, 0xfc0007ff, 0x10000297, 0x1f0000, // Vector Convert Floating-Point Single-Precision to Signed Fraction EVX-form (evfsctsf RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EVFSCTUF, 0xfc0007ff, 0x10000296, 0x1f0000, // Vector Convert Floating-Point Single-Precision to Unsigned Fraction EVX-form (evfsctuf RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFSABS, 0xfc0007ff, 0x100002c4, 0xf800, // Floating-Point Single-Precision Absolute Value EVX-form (efsabs RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EFSNEG, 0xfc0007ff, 0x100002c6, 0xf800, // Floating-Point Single-Precision Negate EVX-form (efsneg RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EFSNABS, 0xfc0007ff, 0x100002c5, 0xf800, // Floating-Point Single-Precision Negative Absolute Value EVX-form (efsnabs RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EFSADD, 0xfc0007ff, 0x100002c0, 0x0, // Floating-Point Single-Precision Add EVX-form (efsadd RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EFSMUL, 0xfc0007ff, 0x100002c8, 0x0, // Floating-Point Single-Precision Multiply EVX-form (efsmul RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EFSSUB, 0xfc0007ff, 0x100002c1, 0x0, // Floating-Point Single-Precision Subtract EVX-form (efssub RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EFSDIV, 0xfc0007ff, 0x100002c9, 0x0, // Floating-Point Single-Precision Divide EVX-form (efsdiv RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EFSCMPGT, 0xfc0007ff, 0x100002cc, 0x600000, // Floating-Point Single-Precision Compare Greater Than EVX-form (efscmpgt BF,RA,RB)
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EFSCMPLT, 0xfc0007ff, 0x100002cd, 0x600000, // Floating-Point Single-Precision Compare Less Than EVX-form (efscmplt BF,RA,RB)
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EFSCMPEQ, 0xfc0007ff, 0x100002ce, 0x600000, // Floating-Point Single-Precision Compare Equal EVX-form (efscmpeq BF,RA,RB)
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EFSTSTGT, 0xfc0007ff, 0x100002dc, 0x600000, // Floating-Point Single-Precision Test Greater Than EVX-form (efststgt BF,RA,RB)
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EFSTSTLT, 0xfc0007ff, 0x100002dd, 0x600000, // Floating-Point Single-Precision Test Less Than EVX-form (efststlt BF,RA,RB)
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EFSTSTEQ, 0xfc0007ff, 0x100002de, 0x600000, // Floating-Point Single-Precision Test Equal EVX-form (efststeq BF,RA,RB)
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EFSCFSI, 0xfc0007ff, 0x100002d1, 0x1f0000, // Convert Floating-Point Single-Precision from Signed Integer EVX-form (efscfsi RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFSCFSF, 0xfc0007ff, 0x100002d3, 0x1f0000, // Convert Floating-Point Single-Precision from Signed Fraction EVX-form (efscfsf RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFSCTSI, 0xfc0007ff, 0x100002d5, 0x1f0000, // Convert Floating-Point Single-Precision to Signed Integer EVX-form (efsctsi RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFSCFUI, 0xfc0007ff, 0x100002d0, 0x1f0000, // Convert Floating-Point Single-Precision from Unsigned Integer EVX-form (efscfui RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFSCFUF, 0xfc0007ff, 0x100002d2, 0x1f0000, // Convert Floating-Point Single-Precision from Unsigned Fraction EVX-form (efscfuf RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFSCTUI, 0xfc0007ff, 0x100002d4, 0x1f0000, // Convert Floating-Point Single-Precision to Unsigned Integer EVX-form (efsctui RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFSCTSIZ, 0xfc0007ff, 0x100002da, 0x1f0000, // Convert Floating-Point Single-Precision to Signed Integer with Round toward Zero EVX-form (efsctsiz RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFSCTSF, 0xfc0007ff, 0x100002d7, 0x1f0000, // Convert Floating-Point Single-Precision to Signed Fraction EVX-form (efsctsf RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFSCTUIZ, 0xfc0007ff, 0x100002d8, 0x1f0000, // Convert Floating-Point Single-Precision to Unsigned Integer with Round toward Zero EVX-form (efsctuiz RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFSCTUF, 0xfc0007ff, 0x100002d6, 0x1f0000, // Convert Floating-Point Single-Precision to Unsigned Fraction EVX-form (efsctuf RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFDABS, 0xfc0007ff, 0x100002e4, 0xf800, // Floating-Point Double-Precision Absolute Value EVX-form (efdabs RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EFDNEG, 0xfc0007ff, 0x100002e6, 0xf800, // Floating-Point Double-Precision Negate EVX-form (efdneg RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EFDNABS, 0xfc0007ff, 0x100002e5, 0xf800, // Floating-Point Double-Precision Negative Absolute Value EVX-form (efdnabs RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EFDADD, 0xfc0007ff, 0x100002e0, 0x0, // Floating-Point Double-Precision Add EVX-form (efdadd RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EFDMUL, 0xfc0007ff, 0x100002e8, 0x0, // Floating-Point Double-Precision Multiply EVX-form (efdmul RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EFDSUB, 0xfc0007ff, 0x100002e1, 0x0, // Floating-Point Double-Precision Subtract EVX-form (efdsub RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EFDDIV, 0xfc0007ff, 0x100002e9, 0x0, // Floating-Point Double-Precision Divide EVX-form (efddiv RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EFDCMPGT, 0xfc0007ff, 0x100002ec, 0x600000, // Floating-Point Double-Precision Compare Greater Than EVX-form (efdcmpgt BF,RA,RB)
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EFDCMPEQ, 0xfc0007ff, 0x100002ee, 0x600000, // Floating-Point Double-Precision Compare Equal EVX-form (efdcmpeq BF,RA,RB)
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EFDCMPLT, 0xfc0007ff, 0x100002ed, 0x600000, // Floating-Point Double-Precision Compare Less Than EVX-form (efdcmplt BF,RA,RB)
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EFDTSTGT, 0xfc0007ff, 0x100002fc, 0x600000, // Floating-Point Double-Precision Test Greater Than EVX-form (efdtstgt BF,RA,RB)
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EFDTSTLT, 0xfc0007ff, 0x100002fd, 0x600000, // Floating-Point Double-Precision Test Less Than EVX-form (efdtstlt BF,RA,RB)
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EFDCFSI, 0xfc0007ff, 0x100002f1, 0x1f0000, // Convert Floating-Point Double-Precision from Signed Integer EVX-form (efdcfsi RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFDTSTEQ, 0xfc0007ff, 0x100002fe, 0x600000, // Floating-Point Double-Precision Test Equal EVX-form (efdtsteq BF,RA,RB)
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EFDCFUI, 0xfc0007ff, 0x100002f0, 0x1f0000, // Convert Floating-Point Double-Precision from Unsigned Integer EVX-form (efdcfui RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFDCFSID, 0xfc0007ff, 0x100002e3, 0x1f0000, // Convert Floating-Point Double-Precision from Signed Integer Doubleword EVX-form (efdcfsid RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFDCFSF, 0xfc0007ff, 0x100002f3, 0x1f0000, // Convert Floating-Point Double-Precision from Signed Fraction EVX-form (efdcfsf RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFDCFUF, 0xfc0007ff, 0x100002f2, 0x1f0000, // Convert Floating-Point Double-Precision from Unsigned Fraction EVX-form (efdcfuf RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFDCFUID, 0xfc0007ff, 0x100002e2, 0x1f0000, // Convert Floating-Point Double-Precision from Unsigned Integer Doubleword EVX-form (efdcfuid RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFDCTSI, 0xfc0007ff, 0x100002f5, 0x1f0000, // Convert Floating-Point Double-Precision to Signed Integer EVX-form (efdctsi RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFDCTUI, 0xfc0007ff, 0x100002f4, 0x1f0000, // Convert Floating-Point Double-Precision to Unsigned Integer EVX-form (efdctui RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFDCTSIDZ, 0xfc0007ff, 0x100002eb, 0x1f0000, // Convert Floating-Point Double-Precision to Signed Integer Doubleword with Round toward Zero EVX-form (efdctsidz RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFDCTUIDZ, 0xfc0007ff, 0x100002ea, 0x1f0000, // Convert Floating-Point Double-Precision to Unsigned Integer Doubleword with Round toward Zero EVX-form (efdctuidz RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFDCTSIZ, 0xfc0007ff, 0x100002fa, 0x1f0000, // Convert Floating-Point Double-Precision to Signed Integer with Round toward Zero EVX-form (efdctsiz RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFDCTSF, 0xfc0007ff, 0x100002f7, 0x1f0000, // Convert Floating-Point Double-Precision to Signed Fraction EVX-form (efdctsf RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFDCTUF, 0xfc0007ff, 0x100002f6, 0x1f0000, // Convert Floating-Point Double-Precision to Unsigned Fraction EVX-form (efdctuf RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFDCTUIZ, 0xfc0007ff, 0x100002f8, 0x1f0000, // Convert Floating-Point Double-Precision to Unsigned Integer with Round toward Zero EVX-form (efdctuiz RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFDCFS, 0xfc0007ff, 0x100002ef, 0x1f0000, // Floating-Point Double-Precision Convert from Single-Precision EVX-form (efdcfs RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFSCFD, 0xfc0007ff, 0x100002cf, 0x1f0000, // Floating-Point Single-Precision Convert from Double-Precision EVX-form (efscfd RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{DLMZB, 0xfc0007ff, 0x7c00009c, 0x0, // Determine Leftmost Zero Byte X-form (dlmzb RA,RS,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{DLMZB_, 0xfc0007ff, 0x7c00009d, 0x0, // Determine Leftmost Zero Byte X-form (dlmzb. RA,RS,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{MACCHW, 0xfc0007ff, 0x10000158, 0x0, // Multiply Accumulate Cross Halfword to Word Modulo Signed XO-form (macchw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACCHW_, 0xfc0007ff, 0x10000159, 0x0, // Multiply Accumulate Cross Halfword to Word Modulo Signed XO-form (macchw. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACCHWO, 0xfc0007ff, 0x10000558, 0x0, // Multiply Accumulate Cross Halfword to Word Modulo Signed XO-form (macchwo RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACCHWO_, 0xfc0007ff, 0x10000559, 0x0, // Multiply Accumulate Cross Halfword to Word Modulo Signed XO-form (macchwo. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACCHWS, 0xfc0007ff, 0x100001d8, 0x0, // Multiply Accumulate Cross Halfword to Word Saturate Signed XO-form (macchws RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACCHWS_, 0xfc0007ff, 0x100001d9, 0x0, // Multiply Accumulate Cross Halfword to Word Saturate Signed XO-form (macchws. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACCHWSO, 0xfc0007ff, 0x100005d8, 0x0, // Multiply Accumulate Cross Halfword to Word Saturate Signed XO-form (macchwso RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACCHWSO_, 0xfc0007ff, 0x100005d9, 0x0, // Multiply Accumulate Cross Halfword to Word Saturate Signed XO-form (macchwso. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACCHWU, 0xfc0007ff, 0x10000118, 0x0, // Multiply Accumulate Cross Halfword to Word Modulo Unsigned XO-form (macchwu RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACCHWU_, 0xfc0007ff, 0x10000119, 0x0, // Multiply Accumulate Cross Halfword to Word Modulo Unsigned XO-form (macchwu. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACCHWUO, 0xfc0007ff, 0x10000518, 0x0, // Multiply Accumulate Cross Halfword to Word Modulo Unsigned XO-form (macchwuo RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACCHWUO_, 0xfc0007ff, 0x10000519, 0x0, // Multiply Accumulate Cross Halfword to Word Modulo Unsigned XO-form (macchwuo. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACCHWSU, 0xfc0007ff, 0x10000198, 0x0, // Multiply Accumulate Cross Halfword to Word Saturate Unsigned XO-form (macchwsu RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACCHWSU_, 0xfc0007ff, 0x10000199, 0x0, // Multiply Accumulate Cross Halfword to Word Saturate Unsigned XO-form (macchwsu. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACCHWSUO, 0xfc0007ff, 0x10000598, 0x0, // Multiply Accumulate Cross Halfword to Word Saturate Unsigned XO-form (macchwsuo RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACCHWSUO_, 0xfc0007ff, 0x10000599, 0x0, // Multiply Accumulate Cross Halfword to Word Saturate Unsigned XO-form (macchwsuo. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACHHW, 0xfc0007ff, 0x10000058, 0x0, // Multiply Accumulate High Halfword to Word Modulo Signed XO-form (machhw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACHHW_, 0xfc0007ff, 0x10000059, 0x0, // Multiply Accumulate High Halfword to Word Modulo Signed XO-form (machhw. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACHHWO, 0xfc0007ff, 0x10000458, 0x0, // Multiply Accumulate High Halfword to Word Modulo Signed XO-form (machhwo RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACHHWO_, 0xfc0007ff, 0x10000459, 0x0, // Multiply Accumulate High Halfword to Word Modulo Signed XO-form (machhwo. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACHHWS, 0xfc0007ff, 0x100000d8, 0x0, // Multiply Accumulate High Halfword to Word Saturate Signed XO-form (machhws RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACHHWS_, 0xfc0007ff, 0x100000d9, 0x0, // Multiply Accumulate High Halfword to Word Saturate Signed XO-form (machhws. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACHHWSO, 0xfc0007ff, 0x100004d8, 0x0, // Multiply Accumulate High Halfword to Word Saturate Signed XO-form (machhwso RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACHHWSO_, 0xfc0007ff, 0x100004d9, 0x0, // Multiply Accumulate High Halfword to Word Saturate Signed XO-form (machhwso. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACHHWU, 0xfc0007ff, 0x10000018, 0x0, // Multiply Accumulate High Halfword to Word Modulo Unsigned XO-form (machhwu RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACHHWU_, 0xfc0007ff, 0x10000019, 0x0, // Multiply Accumulate High Halfword to Word Modulo Unsigned XO-form (machhwu. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACHHWUO, 0xfc0007ff, 0x10000418, 0x0, // Multiply Accumulate High Halfword to Word Modulo Unsigned XO-form (machhwuo RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACHHWUO_, 0xfc0007ff, 0x10000419, 0x0, // Multiply Accumulate High Halfword to Word Modulo Unsigned XO-form (machhwuo. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACHHWSU, 0xfc0007ff, 0x10000098, 0x0, // Multiply Accumulate High Halfword to Word Saturate Unsigned XO-form (machhwsu RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACHHWSU_, 0xfc0007ff, 0x10000099, 0x0, // Multiply Accumulate High Halfword to Word Saturate Unsigned XO-form (machhwsu. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACHHWSUO, 0xfc0007ff, 0x10000498, 0x0, // Multiply Accumulate High Halfword to Word Saturate Unsigned XO-form (machhwsuo RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACHHWSUO_, 0xfc0007ff, 0x10000499, 0x0, // Multiply Accumulate High Halfword to Word Saturate Unsigned XO-form (machhwsuo. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACLHW, 0xfc0007ff, 0x10000358, 0x0, // Multiply Accumulate Low Halfword to Word Modulo Signed XO-form (maclhw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACLHW_, 0xfc0007ff, 0x10000359, 0x0, // Multiply Accumulate Low Halfword to Word Modulo Signed XO-form (maclhw. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACLHWO, 0xfc0007ff, 0x10000758, 0x0, // Multiply Accumulate Low Halfword to Word Modulo Signed XO-form (maclhwo RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACLHWO_, 0xfc0007ff, 0x10000759, 0x0, // Multiply Accumulate Low Halfword to Word Modulo Signed XO-form (maclhwo. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACLHWS, 0xfc0007ff, 0x100003d8, 0x0, // Multiply Accumulate Low Halfword to Word Saturate Signed XO-form (maclhws RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACLHWS_, 0xfc0007ff, 0x100003d9, 0x0, // Multiply Accumulate Low Halfword to Word Saturate Signed XO-form (maclhws. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACLHWSO, 0xfc0007ff, 0x100007d8, 0x0, // Multiply Accumulate Low Halfword to Word Saturate Signed XO-form (maclhwso RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACLHWSO_, 0xfc0007ff, 0x100007d9, 0x0, // Multiply Accumulate Low Halfword to Word Saturate Signed XO-form (maclhwso. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACLHWU, 0xfc0007ff, 0x10000318, 0x0, // Multiply Accumulate Low Halfword to Word Modulo Unsigned XO-form (maclhwu RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACLHWU_, 0xfc0007ff, 0x10000319, 0x0, // Multiply Accumulate Low Halfword to Word Modulo Unsigned XO-form (maclhwu. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACLHWUO, 0xfc0007ff, 0x10000718, 0x0, // Multiply Accumulate Low Halfword to Word Modulo Unsigned XO-form (maclhwuo RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACLHWUO_, 0xfc0007ff, 0x10000719, 0x0, // Multiply Accumulate Low Halfword to Word Modulo Unsigned XO-form (maclhwuo. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULCHW, 0xfc0007ff, 0x10000150, 0x0, // Multiply Cross Halfword to Word Signed X-form (mulchw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULCHW_, 0xfc0007ff, 0x10000151, 0x0, // Multiply Cross Halfword to Word Signed X-form (mulchw. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACLHWSU, 0xfc0007ff, 0x10000398, 0x0, // Multiply Accumulate Low Halfword to Word Saturate Unsigned XO-form (maclhwsu RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACLHWSU_, 0xfc0007ff, 0x10000399, 0x0, // Multiply Accumulate Low Halfword to Word Saturate Unsigned XO-form (maclhwsu. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACLHWSUO, 0xfc0007ff, 0x10000798, 0x0, // Multiply Accumulate Low Halfword to Word Saturate Unsigned XO-form (maclhwsuo RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACLHWSUO_, 0xfc0007ff, 0x10000799, 0x0, // Multiply Accumulate Low Halfword to Word Saturate Unsigned XO-form (maclhwsuo. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULCHWU, 0xfc0007ff, 0x10000110, 0x0, // Multiply Cross Halfword to Word Unsigned X-form (mulchwu RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULCHWU_, 0xfc0007ff, 0x10000111, 0x0, // Multiply Cross Halfword to Word Unsigned X-form (mulchwu. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULHHW, 0xfc0007ff, 0x10000050, 0x0, // Multiply High Halfword to Word Signed X-form (mulhhw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULHHW_, 0xfc0007ff, 0x10000051, 0x0, // Multiply High Halfword to Word Signed X-form (mulhhw. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULLHW, 0xfc0007ff, 0x10000350, 0x0, // Multiply Low Halfword to Word Signed X-form (mullhw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULLHW_, 0xfc0007ff, 0x10000351, 0x0, // Multiply Low Halfword to Word Signed X-form (mullhw. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULHHWU, 0xfc0007ff, 0x10000010, 0x0, // Multiply High Halfword to Word Unsigned X-form (mulhhwu RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULHHWU_, 0xfc0007ff, 0x10000011, 0x0, // Multiply High Halfword to Word Unsigned X-form (mulhhwu. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULLHWU, 0xfc0007ff, 0x10000310, 0x0, // Multiply Low Halfword to Word Unsigned X-form (mullhwu RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULLHWU_, 0xfc0007ff, 0x10000311, 0x0, // Multiply Low Halfword to Word Unsigned X-form (mullhwu. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACCHW, 0xfc0007ff, 0x1000015c, 0x0, // Negative Multiply Accumulate Cross Halfword to Word Modulo Signed XO-form (nmacchw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACCHW_, 0xfc0007ff, 0x1000015d, 0x0, // Negative Multiply Accumulate Cross Halfword to Word Modulo Signed XO-form (nmacchw. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACCHWO, 0xfc0007ff, 0x1000055c, 0x0, // Negative Multiply Accumulate Cross Halfword to Word Modulo Signed XO-form (nmacchwo RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACCHWO_, 0xfc0007ff, 0x1000055d, 0x0, // Negative Multiply Accumulate Cross Halfword to Word Modulo Signed XO-form (nmacchwo. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACCHWS, 0xfc0007ff, 0x100001dc, 0x0, // Negative Multiply Accumulate Cross Halfword to Word Saturate Signed XO-form (nmacchws RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACCHWS_, 0xfc0007ff, 0x100001dd, 0x0, // Negative Multiply Accumulate Cross Halfword to Word Saturate Signed XO-form (nmacchws. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACCHWSO, 0xfc0007ff, 0x100005dc, 0x0, // Negative Multiply Accumulate Cross Halfword to Word Saturate Signed XO-form (nmacchwso RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACCHWSO_, 0xfc0007ff, 0x100005dd, 0x0, // Negative Multiply Accumulate Cross Halfword to Word Saturate Signed XO-form (nmacchwso. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACHHW, 0xfc0007ff, 0x1000005c, 0x0, // Negative Multiply Accumulate High Halfword to Word Modulo Signed XO-form (nmachhw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACHHW_, 0xfc0007ff, 0x1000005d, 0x0, // Negative Multiply Accumulate High Halfword to Word Modulo Signed XO-form (nmachhw. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACHHWO, 0xfc0007ff, 0x1000045c, 0x0, // Negative Multiply Accumulate High Halfword to Word Modulo Signed XO-form (nmachhwo RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACHHWO_, 0xfc0007ff, 0x1000045d, 0x0, // Negative Multiply Accumulate High Halfword to Word Modulo Signed XO-form (nmachhwo. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACHHWS, 0xfc0007ff, 0x100000dc, 0x0, // Negative Multiply Accumulate High Halfword to Word Saturate Signed XO-form (nmachhws RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACHHWS_, 0xfc0007ff, 0x100000dd, 0x0, // Negative Multiply Accumulate High Halfword to Word Saturate Signed XO-form (nmachhws. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACHHWSO, 0xfc0007ff, 0x100004dc, 0x0, // Negative Multiply Accumulate High Halfword to Word Saturate Signed XO-form (nmachhwso RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACHHWSO_, 0xfc0007ff, 0x100004dd, 0x0, // Negative Multiply Accumulate High Halfword to Word Saturate Signed XO-form (nmachhwso. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACLHW, 0xfc0007ff, 0x1000035c, 0x0, // Negative Multiply Accumulate Low Halfword to Word Modulo Signed XO-form (nmaclhw RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACLHW_, 0xfc0007ff, 0x1000035d, 0x0, // Negative Multiply Accumulate Low Halfword to Word Modulo Signed XO-form (nmaclhw. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACLHWO, 0xfc0007ff, 0x1000075c, 0x0, // Negative Multiply Accumulate Low Halfword to Word Modulo Signed XO-form (nmaclhwo RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACLHWO_, 0xfc0007ff, 0x1000075d, 0x0, // Negative Multiply Accumulate Low Halfword to Word Modulo Signed XO-form (nmaclhwo. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACLHWS, 0xfc0007ff, 0x100003dc, 0x0, // Negative Multiply Accumulate Low Halfword to Word Saturate Signed XO-form (nmaclhws RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACLHWS_, 0xfc0007ff, 0x100003dd, 0x0, // Negative Multiply Accumulate Low Halfword to Word Saturate Signed XO-form (nmaclhws. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACLHWSO, 0xfc0007ff, 0x100007dc, 0x0, // Negative Multiply Accumulate Low Halfword to Word Saturate Signed XO-form (nmaclhwso RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACLHWSO_, 0xfc0007ff, 0x100007dd, 0x0, // Negative Multiply Accumulate Low Halfword to Word Saturate Signed XO-form (nmaclhwso. RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ICBI, 0xfc0007fe, 0x7c0007ac, 0x3e00001, // Instruction Cache Block Invalidate X-form (icbi RA,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_16_20}},
	{ICBT, 0xfc0007fe, 0x7c00002c, 0x2000001, // Instruction Cache Block Touch X-form (icbt CT, RA, RB)
		[5]*argField{ap_ImmUnsigned_7_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DCBA, 0xfc0007fe, 0x7c0005ec, 0x3e00001, // Data Cache Block Allocate X-form (dcba RA,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_16_20}},
	{DCBT, 0xfc0007fe, 0x7c00022c, 0x1, // Data Cache Block Touch X-form (dcbt RA,RB,TH)
		[5]*argField{ap_Reg_11_15, ap_Reg_16_20, ap_ImmUnsigned_6_10}},
	{DCBT, 0xfc0007fe, 0x7c00022c, 0x1, // Data Cache Block Touch X-form (dcbt TH,RA,RB)
		[5]*argField{ap_ImmUnsigned_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DCBTST, 0xfc0007fe, 0x7c0001ec, 0x1, // Data Cache Block Touch for Store X-form (dcbtst RA,RB,TH)
		[5]*argField{ap_Reg_11_15, ap_Reg_16_20, ap_ImmUnsigned_6_10}},
	{DCBTST, 0xfc0007fe, 0x7c0001ec, 0x1, // Data Cache Block Touch for Store X-form (dcbtst TH,RA,RB)
		[5]*argField{ap_ImmUnsigned_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DCBZ, 0xfc0007fe, 0x7c0007ec, 0x3e00001, // Data Cache Block set to Zero X-form (dcbz RA,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_16_20}},
	{DCBST, 0xfc0007fe, 0x7c00006c, 0x3e00001, // Data Cache Block Store X-form (dcbst RA,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_16_20}},
	{DCBF, 0xfc0007fe, 0x7c0000ac, 0x3800001, // Data Cache Block Flush X-form (dcbf RA,RB,L)
		[5]*argField{ap_Reg_11_15, ap_Reg_16_20, ap_ImmUnsigned_9_10}},
	{ISYNC, 0xfc0007fe, 0x4c00012c, 0x3fff801, // Instruction Synchronize XL-form (isync)
		[5]*argField{}},
	{LBARX, 0xfc0007ff, 0x7c000068, 0x0, // Load Byte And Reserve Indexed X-form [Category: Phased-In] (lbarx RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LBARX, 0xfc0007fe, 0x7c000068, 0x0, // Load Byte And Reserve Indexed X-form [Category: Phased-In] (lbarx RT,RA,RB,EH)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20, ap_ImmUnsigned_31_31}},
	{LHARX, 0xfc0007ff, 0x7c0000e8, 0x0, // Load Halfword And Reserve Indexed X-form [Category: Phased-In] (lharx RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LHARX, 0xfc0007fe, 0x7c0000e8, 0x0, // Load Halfword And Reserve Indexed X-form [Category: Phased-In] (lharx RT,RA,RB,EH)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20, ap_ImmUnsigned_31_31}},
	{LWARX, 0xfc0007ff, 0x7c000028, 0x0, // Load Word And Reserve Indexed X-form (lwarx RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LWARX, 0xfc0007ff, 0x7c000028, 0x0, // Load Word And Reserve Indexed X-form (lwarx RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LWARX, 0xfc0007fe, 0x7c000028, 0x0, // Load Word And Reserve Indexed X-form (lwarx RT,RA,RB,EH)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20, ap_ImmUnsigned_31_31}},
	{STBCX_, 0xfc0007ff, 0x7c00056d, 0x0, // Store Byte Conditional Indexed X-form [Category: Phased-In] (stbcx. RS,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STHCX_, 0xfc0007ff, 0x7c0005ad, 0x0, // Store Halfword Conditional Indexed X-form [Category: Phased-In] (sthcx. RS,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STWCX_, 0xfc0007ff, 0x7c00012d, 0x0, // Store Word Conditional Indexed X-form (stwcx. RS,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LDARX, 0xfc0007ff, 0x7c0000a8, 0x0, // Load Doubleword And Reserve Indexed X-form (ldarx RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LDARX, 0xfc0007fe, 0x7c0000a8, 0x0, // Load Doubleword And Reserve Indexed X-form (ldarx RT,RA,RB,EH)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20, ap_ImmUnsigned_31_31}},
	{STDCX_, 0xfc0007ff, 0x7c0001ad, 0x0, // Store Doubleword Conditional Indexed X-form (stdcx. RS,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LQARX, 0xfc0007ff, 0x7c000228, 0x0, // Load Quadword And Reserve Indexed X-form (lqarx RTp,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LQARX, 0xfc0007fe, 0x7c000228, 0x0, // Load Quadword And Reserve Indexed X-form (lqarx RTp,RA,RB,EH)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20, ap_ImmUnsigned_31_31}},
	{STQCX_, 0xfc0007ff, 0x7c00016d, 0x0, // Store Quadword Conditional Indexed X-form (stqcx. RSp,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{SYNC, 0xfc0007fe, 0x7c0004ac, 0x390f801, // Synchronize X-form (sync L, E)
		[5]*argField{ap_ImmUnsigned_9_10, ap_ImmUnsigned_12_15}},
	{EIEIO, 0xfc0007fe, 0x7c0006ac, 0x3fff801, // Enforce In-order Execution of I/O X-form (eieio)
		[5]*argField{}},
	{MBAR, 0xfc0007fe, 0x7c0006ac, 0x1ff801, // Memory Barrier X-form (mbar MO)
		[5]*argField{ap_ImmUnsigned_6_10}},
	{WAIT, 0xfc0007fe, 0x7c00007c, 0x39ff801, // Wait X-form (wait WC)
		[5]*argField{ap_ImmUnsigned_9_10}},
	{TBEGIN_, 0xfc0007ff, 0x7c00051d, 0x1dff800, // Transaction Begin X-form (tbegin. R)
		[5]*argField{ap_ImmUnsigned_10_10}},
	{TEND_, 0xfc0007ff, 0x7c00055d, 0x1fff800, // Transaction End X-form (tend. A)
		[5]*argField{ap_ImmUnsigned_6_6}},
	{TABORT_, 0xfc0007ff, 0x7c00071d, 0x3e0f800, // Transaction Abort X-form (tabort. RA)
		[5]*argField{ap_Reg_11_15}},
	{TABORTWC_, 0xfc0007ff, 0x7c00061d, 0x0, // Transaction Abort Word Conditional X-form (tabortwc. TO,RA,RB)
		[5]*argField{ap_ImmUnsigned_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{TABORTWCI_, 0xfc0007ff, 0x7c00069d, 0x0, // Transaction Abort Word Conditional Immediate X-form (tabortwci. TO,RA,SI)
		[5]*argField{ap_ImmUnsigned_6_10, ap_Reg_11_15, ap_ImmSigned_16_20}},
	{TABORTDC_, 0xfc0007ff, 0x7c00065d, 0x0, // Transaction Abort Doubleword Conditional X-form (tabortdc. TO,RA,RB)
		[5]*argField{ap_ImmUnsigned_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{TABORTDCI_, 0xfc0007ff, 0x7c0006dd, 0x0, // Transaction Abort Doubleword Conditional Immediate X-form (tabortdci. TO,RA, SI)
		[5]*argField{ap_ImmUnsigned_6_10, ap_Reg_11_15, ap_ImmSigned_16_20}},
	{TSR_, 0xfc0007ff, 0x7c0005dd, 0x3dff800, // Transaction Suspend or Resume X-form (tsr. L)
		[5]*argField{ap_ImmUnsigned_10_10}},
	{TCHECK, 0xfc0007fe, 0x7c00059c, 0x7ff801, // Transaction Check X-form (tcheck BF)
		[5]*argField{ap_CondRegField_6_8}},
	{MFTB, 0xfc0007fe, 0x7c0002e6, 0x1, // Move From Time Base XFX-form (mftb RT,TBR)
		[5]*argField{ap_Reg_6_10, ap_SpReg_16_20_11_15}},
	{RFEBB, 0xfc0007fe, 0x4c000124, 0x3fff001, // Return from Event-Based Branch XL-form (rfebb S)
		[5]*argField{ap_ImmUnsigned_20_20}},
	{LBDX, 0xfc0007fe, 0x7c000406, 0x1, // Load Byte with Decoration Indexed X-form (lbdx RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LHDX, 0xfc0007fe, 0x7c000446, 0x1, // Load Halfword with Decoration Indexed X-form (lhdx RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LWDX, 0xfc0007fe, 0x7c000486, 0x1, // Load Word with Decoration Indexed X-form (lwdx RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LDDX, 0xfc0007fe, 0x7c0004c6, 0x1, // Load Doubleword with Decoration Indexed X-form (lddx RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LFDDX, 0xfc0007fe, 0x7c000646, 0x1, // Load Floating Doubleword with Decoration Indexed X-form (lfddx FRT,RA,RB)
		[5]*argField{ap_FPReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STBDX, 0xfc0007fe, 0x7c000506, 0x1, // Store Byte with Decoration Indexed X-form (stbdx RS,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STHDX, 0xfc0007fe, 0x7c000546, 0x1, // Store Halfword with Decoration Indexed X-form (sthdx RS,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STWDX, 0xfc0007fe, 0x7c000586, 0x1, // Store Word with Decoration Indexed X-form (stwdx RS,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STDDX, 0xfc0007fe, 0x7c0005c6, 0x1, // Store Doubleword with Decoration Indexed X-form (stddx RS,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STFDDX, 0xfc0007fe, 0x7c000746, 0x1, // Store Floating Doubleword with Decoration Indexed X-form (stfddx FRS,RA,RB)
		[5]*argField{ap_FPReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DSN, 0xfc0007fe, 0x7c0003c6, 0x3e00001, // Decorated Storage Notify X-form (dsn RA,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_16_20}},
	{ECIWX, 0xfc0007fe, 0x7c00026c, 0x1, // External Control In Word Indexed X-form (eciwx RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ECOWX, 0xfc0007fe, 0x7c00036c, 0x1, // External Control Out Word Indexed X-form (ecowx RS,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{SC, 0xfc000002, 0x44000002, 0x3fff01d, // System Call SC-form (sc LEV)
		[5]*argField{ap_ImmUnsigned_20_26}},
	{RFID, 0xfc0007fe, 0x4c000024, 0x3fff801, // Return From Interrupt Doubleword XL-form (rfid)
		[5]*argField{}},
	{HRFID, 0xfc0007fe, 0x4c000224, 0x3fff801, // Hypervisor Return From Interrupt Doubleword XL-form (hrfid)
		[5]*argField{}},
	{DOZE, 0xfc0007fe, 0x4c000324, 0x3fff801, // Doze XL-form (doze)
		[5]*argField{}},
	{NAP, 0xfc0007fe, 0x4c000364, 0x3fff801, // Nap XL-form (nap)
		[5]*argField{}},
	{SLEEP, 0xfc0007fe, 0x4c0003a4, 0x3fff801, // Sleep XL-form (sleep)
		[5]*argField{}},
	{RVWINKLE, 0xfc0007fe, 0x4c0003e4, 0x3fff801, // Rip Van Winkle XL-form (rvwinkle)
		[5]*argField{}},
	{LBZCIX, 0xfc0007fe, 0x7c0006aa, 0x1, // Load Byte and Zero Caching Inhibited Indexed X-form (lbzcix RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LWZCIX, 0xfc0007fe, 0x7c00062a, 0x1, // Load Word and Zero Caching Inhibited Indexed X-form (lwzcix RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LHZCIX, 0xfc0007fe, 0x7c00066a, 0x1, // Load Halfword and Zero Caching Inhibited Indexed X-form (lhzcix RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LDCIX, 0xfc0007fe, 0x7c0006ea, 0x1, // Load Doubleword Caching Inhibited Indexed X-form (ldcix RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STBCIX, 0xfc0007fe, 0x7c0007aa, 0x1, // Store Byte Caching Inhibited Indexed X-form (stbcix RS,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STWCIX, 0xfc0007fe, 0x7c00072a, 0x1, // Store Word Caching Inhibited Indexed X-form (stwcix RS,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STHCIX, 0xfc0007fe, 0x7c00076a, 0x1, // Store Halfword Caching Inhibited Indexed X-form (sthcix RS,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STDCIX, 0xfc0007fe, 0x7c0007ea, 0x1, // Store Doubleword Caching Inhibited Indexed X-form (stdcix RS,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{TRECLAIM_, 0xfc0007ff, 0x7c00075d, 0x3e0f800, // Transaction Reclaim X-form (treclaim. RA)
		[5]*argField{ap_Reg_11_15}},
	{TRECHKPT_, 0xfc0007ff, 0x7c0007dd, 0x3fff800, // Transaction Recheckpoint X-form (trechkpt.)
		[5]*argField{}},
	{MTSPR, 0xfc0007fe, 0x7c0003a6, 0x1, // Move To Special Purpose Register XFX-form (mtspr SPR,RS)
		[5]*argField{ap_SpReg_16_20_11_15, ap_Reg_6_10}},
	{MFSPR, 0xfc0007fe, 0x7c0002a6, 0x1, // Move From Special Purpose Register XFX-form (mfspr RT,SPR)
		[5]*argField{ap_Reg_6_10, ap_SpReg_16_20_11_15}},
	{MTMSR, 0xfc0007fe, 0x7c000124, 0x1ef801, // Move To Machine State Register X-form (mtmsr RS,L)
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_15_15}},
	{MTMSRD, 0xfc0007fe, 0x7c000164, 0x1ef801, // Move To Machine State Register Doubleword X-form (mtmsrd RS,L)
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_15_15}},
	{MFMSR, 0xfc0007fe, 0x7c0000a6, 0x1ff801, // Move From Machine State Register X-form (mfmsr RT)
		[5]*argField{ap_Reg_6_10}},
	{SLBIE, 0xfc0007fe, 0x7c000364, 0x3ff0001, // SLB Invalidate Entry X-form (slbie RB)
		[5]*argField{ap_Reg_16_20}},
	{SLBIA, 0xfc0007fe, 0x7c0003e4, 0x31ff801, // SLB Invalidate All X-form (slbia IH)
		[5]*argField{ap_ImmUnsigned_8_10}},
	{SLBMTE, 0xfc0007fe, 0x7c000324, 0x1f0001, // SLB Move To Entry X-form (slbmte RS,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{SLBMFEV, 0xfc0007fe, 0x7c0006a6, 0x1f0001, // SLB Move From Entry VSID X-form (slbmfev RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{SLBMFEE, 0xfc0007fe, 0x7c000726, 0x1f0001, // SLB Move From Entry ESID X-form (slbmfee RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{SLBFEE_, 0xfc0007ff, 0x7c0007a7, 0x1f0000, // SLB Find Entry ESID X-form (slbfee. RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{MTSR, 0xfc0007fe, 0x7c0001a4, 0x10f801, // Move To Segment Register X-form (mtsr SR,RS)
		[5]*argField{ap_SpReg_12_15, ap_Reg_6_10}},
	{MTSRIN, 0xfc0007fe, 0x7c0001e4, 0x1f0001, // Move To Segment Register Indirect X-form (mtsrin RS,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{MFSR, 0xfc0007fe, 0x7c0004a6, 0x10f801, // Move From Segment Register X-form (mfsr RT,SR)
		[5]*argField{ap_Reg_6_10, ap_SpReg_12_15}},
	{MFSRIN, 0xfc0007fe, 0x7c000526, 0x1f0001, // Move From Segment Register Indirect X-form (mfsrin RT,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{TLBIE, 0xfc0007fe, 0x7c000264, 0x1f0001, // TLB Invalidate Entry X-form (tlbie RB,RS)
		[5]*argField{ap_Reg_16_20, ap_Reg_6_10}},
	{TLBIEL, 0xfc0007fe, 0x7c000224, 0x3ff0001, // TLB Invalidate Entry Local X-form (tlbiel RB)
		[5]*argField{ap_Reg_16_20}},
	{TLBIA, 0xfc0007fe, 0x7c0002e4, 0x3fff801, // TLB Invalidate All X-form (tlbia)
		[5]*argField{}},
	{TLBSYNC, 0xfc0007fe, 0x7c00046c, 0x3fff801, // TLB Synchronize X-form (tlbsync)
		[5]*argField{}},
	{MSGSND, 0xfc0007fe, 0x7c00019c, 0x3ff0001, // Message Send X-form (msgsnd RB)
		[5]*argField{ap_Reg_16_20}},
	{MSGCLR, 0xfc0007fe, 0x7c0001dc, 0x3ff0001, // Message Clear X-form (msgclr RB)
		[5]*argField{ap_Reg_16_20}},
	{MSGSNDP, 0xfc0007fe, 0x7c00011c, 0x3ff0001, // Message Send Privileged X-form (msgsndp RB)
		[5]*argField{ap_Reg_16_20}},
	{MSGCLRP, 0xfc0007fe, 0x7c00015c, 0x3ff0001, // Message Clear Privileged X-form (msgclrp RB)
		[5]*argField{ap_Reg_16_20}},
	{MTTMR, 0xfc0007fe, 0x7c0003dc, 0x1, // Move To Thread Management Register XFX-form (mttmr TMR,RS)
		[5]*argField{ap_SpReg_16_20_11_15, ap_Reg_6_10}},
	{SC, 0xfc000002, 0x44000002, 0x3fffffd, // System Call SC-form (sc)
		[5]*argField{}},
	{RFI, 0xfc0007fe, 0x4c000064, 0x3fff801, // Return From Interrupt XL-form (rfi)
		[5]*argField{}},
	{RFCI, 0xfc0007fe, 0x4c000066, 0x3fff801, // Return From Critical Interrupt XL-form (rfci)
		[5]*argField{}},
	{RFDI, 0xfc0007fe, 0x4c00004e, 0x3fff801, // Return From Debug Interrupt X-form (rfdi)
		[5]*argField{}},
	{RFMCI, 0xfc0007fe, 0x4c00004c, 0x3fff801, // Return From Machine Check Interrupt XL-form (rfmci)
		[5]*argField{}},
	{RFGI, 0xfc0007fe, 0x4c0000cc, 0x3fff801, // Return From Guest Interrupt XL-form (rfgi)
		[5]*argField{}},
	{EHPRIV, 0xfc0007fe, 0x7c00021c, 0x1, // Embedded Hypervisor Privilege XL-form (ehpriv OC)
		[5]*argField{ap_ImmUnsigned_6_20}},
	{MTSPR, 0xfc0007fe, 0x7c0003a6, 0x1, // Move To Special Purpose Register XFX-form (mtspr SPR,RS)
		[5]*argField{ap_SpReg_16_20_11_15, ap_Reg_6_10}},
	{MFSPR, 0xfc0007fe, 0x7c0002a6, 0x1, // Move From Special Purpose Register XFX-form (mfspr RT,SPR)
		[5]*argField{ap_Reg_6_10, ap_SpReg_16_20_11_15}},
	{MTDCR, 0xfc0007fe, 0x7c000386, 0x1, // Move To Device Control Register XFX-form (mtdcr DCRN,RS)
		[5]*argField{ap_SpReg_16_20_11_15, ap_Reg_6_10}},
	{MTDCRX, 0xfc0007fe, 0x7c000306, 0xf801, // Move To Device Control Register Indexed X-form (mtdcrx RA,RS)
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10}},
	{MFDCR, 0xfc0007fe, 0x7c000286, 0x1, // Move From Device Control Register XFX-form (mfdcr RT,DCRN)
		[5]*argField{ap_Reg_6_10, ap_SpReg_16_20_11_15}},
	{MFDCRX, 0xfc0007fe, 0x7c000206, 0xf801, // Move From Device Control Register Indexed X-form (mfdcrx RT,RA)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{MTMSR, 0xfc0007fe, 0x7c000124, 0x1ff801, // Move To Machine State Register X-form (mtmsr RS)
		[5]*argField{ap_Reg_6_10}},
	{MFMSR, 0xfc0007fe, 0x7c0000a6, 0x1ff801, // Move From Machine State Register X-form (mfmsr RT)
		[5]*argField{ap_Reg_6_10}},
	{WRTEE, 0xfc0007fe, 0x7c000106, 0x1ff801, // Write MSR External Enable X-form (wrtee RS)
		[5]*argField{ap_Reg_6_10}},
	{WRTEEI, 0xfc0007fe, 0x7c000146, 0x3ff7801, // Write MSR External Enable Immediate X-form (wrteei E)
		[5]*argField{ap_ImmUnsigned_16_16}},
	{LBEPX, 0xfc0007fe, 0x7c0000be, 0x1, // Load Byte by External Process ID Indexed X-form (lbepx RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LHEPX, 0xfc0007fe, 0x7c00023e, 0x1, // Load Halfword by External Process ID Indexed X-form (lhepx RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LWEPX, 0xfc0007fe, 0x7c00003e, 0x1, // Load Word by External Process ID Indexed X-form (lwepx RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LDEPX, 0xfc0007fe, 0x7c00003a, 0x1, // Load Doubleword by External Process ID Indexed X-form (ldepx RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STBEPX, 0xfc0007fe, 0x7c0001be, 0x1, // Store Byte by External Process ID Indexed X-form (stbepx RS,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STHEPX, 0xfc0007fe, 0x7c00033e, 0x1, // Store Halfword by External Process ID Indexed X-form (sthepx RS,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STWEPX, 0xfc0007fe, 0x7c00013e, 0x1, // Store Word by External Process ID Indexed X-form (stwepx RS,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STDEPX, 0xfc0007fe, 0x7c00013a, 0x1, // Store Doubleword by External Process ID Indexed X-form (stdepx RS,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DCBSTEP, 0xfc0007fe, 0x7c00007e, 0x3e00001, // Data Cache Block Store by External PID X-form (dcbstep RA,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_16_20}},
	{DCBTEP, 0xfc0007fe, 0x7c00027e, 0x1, // Data Cache Block Touch by External PID X-form (dcbtep TH,RA,RB)
		[5]*argField{ap_ImmUnsigned_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DCBFEP, 0xfc0007fe, 0x7c0000fe, 0x3800001, // Data Cache Block Flush by External PID X-form (dcbfep RA,RB,L)
		[5]*argField{ap_Reg_11_15, ap_Reg_16_20, ap_ImmUnsigned_9_10}},
	{DCBTSTEP, 0xfc0007fe, 0x7c0001fe, 0x1, // Data Cache Block Touch for Store by External PID X-form (dcbtstep TH,RA,RB)
		[5]*argField{ap_ImmUnsigned_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ICBIEP, 0xfc0007fe, 0x7c0007be, 0x3e00001, // Instruction Cache Block Invalidate by External PID X-form (icbiep RA,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_16_20}},
	{DCBZEP, 0xfc0007fe, 0x7c0007fe, 0x3e00001, // Data Cache Block set to Zero by External PID X-form (dcbzep RA,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_16_20}},
	{LFDEPX, 0xfc0007fe, 0x7c0004be, 0x1, // Load Floating-Point Double by External Process ID Indexed X-form (lfdepx FRT,RA,RB)
		[5]*argField{ap_FPReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STFDEPX, 0xfc0007fe, 0x7c0005be, 0x1, // Store Floating-Point Double by External Process ID Indexed X-form (stfdepx FRS,RA,RB)
		[5]*argField{ap_FPReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVLDDEPX, 0xfc0007fe, 0x7c00063e, 0x1, // Vector Load Doubleword into Doubleword by External Process ID Indexed EVX-form (evlddepx RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVSTDDEPX, 0xfc0007fe, 0x7c00073e, 0x1, // Vector Store Doubleword into Doubleword by External Process ID Indexed EVX-form (evstddepx RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LVEPX, 0xfc0007fe, 0x7c00024e, 0x1, // Load Vector by External Process ID Indexed X-form (lvepx VRT,RA,RB)
		[5]*argField{ap_VecReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LVEPXL, 0xfc0007fe, 0x7c00020e, 0x1, // Load Vector by External Process ID Indexed LRU X-form (lvepxl VRT,RA,RB)
		[5]*argField{ap_VecReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STVEPX, 0xfc0007fe, 0x7c00064e, 0x1, // Store Vector by External Process ID Indexed X-form (stvepx VRS,RA,RB)
		[5]*argField{ap_VecReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STVEPXL, 0xfc0007fe, 0x7c00060e, 0x1, // Store Vector by External Process ID Indexed LRU X-form (stvepxl VRS,RA,RB)
		[5]*argField{ap_VecReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DCBI, 0xfc0007fe, 0x7c0003ac, 0x3e00001, // Data Cache Block Invalidate X-form (dcbi RA,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_16_20}},
	{DCBLQ_, 0xfc0007ff, 0x7c00034d, 0x2000000, // Data Cache Block Lock Query X-form (dcblq. CT,RA,RB)
		[5]*argField{ap_ImmUnsigned_7_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ICBLQ_, 0xfc0007ff, 0x7c00018d, 0x2000000, // Instruction Cache Block Lock Query X-form (icblq. CT,RA,RB)
		[5]*argField{ap_ImmUnsigned_7_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DCBTLS, 0xfc0007fe, 0x7c00014c, 0x2000001, // Data Cache Block Touch and Lock Set X-form (dcbtls CT,RA,RB)
		[5]*argField{ap_ImmUnsigned_7_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DCBTSTLS, 0xfc0007fe, 0x7c00010c, 0x2000001, // Data Cache Block Touch for Store and Lock Set X-form (dcbtstls CT,RA,RB)
		[5]*argField{ap_ImmUnsigned_7_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ICBTLS, 0xfc0007fe, 0x7c0003cc, 0x2000001, // Instruction Cache Block Touch and Lock Set X-form (icbtls CT,RA,RB)
		[5]*argField{ap_ImmUnsigned_7_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ICBLC, 0xfc0007fe, 0x7c0001cc, 0x2000001, // Instruction Cache Block Lock Clear X-form (icblc CT,RA,RB)
		[5]*argField{ap_ImmUnsigned_7_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DCBLC, 0xfc0007fe, 0x7c00030c, 0x2000001, // Data Cache Block Lock Clear X-form (dcblc CT,RA,RB)
		[5]*argField{ap_ImmUnsigned_7_10, ap_Reg_11_15, ap_Reg_16_20}},
	{TLBIVAX, 0xfc0007fe, 0x7c000624, 0x3e00001, // TLB Invalidate Virtual Address Indexed X-form (tlbivax RA,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_16_20}},
	{TLBILX, 0xfc0007fe, 0x7c000024, 0x3800001, // TLB Invalidate Local Indexed X-form (tlbilx RA,RB])
		[5]*argField{ap_Reg_11_15, ap_Reg_16_20}},
	{TLBSX, 0xfc0007fe, 0x7c000724, 0x3e00001, // TLB Search Indexed X-form (tlbsx RA,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_16_20}},
	{TLBSRX_, 0xfc0007ff, 0x7c0006a5, 0x3e00000, // TLB Search and Reserve Indexed X-form (tlbsrx. RA,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_16_20}},
	{TLBRE, 0xfc0007fe, 0x7c000764, 0x3fff801, // TLB Read Entry X-form (tlbre)
		[5]*argField{}},
	{TLBSYNC, 0xfc0007fe, 0x7c00046c, 0x3fff801, // TLB Synchronize X-form (tlbsync)
		[5]*argField{}},
	{TLBWE, 0xfc0007fe, 0x7c0007a4, 0x3fff801, // TLB Write Entry X-form (tlbwe)
		[5]*argField{}},
	{DNH, 0xfc0007fe, 0x4c00018c, 0x1, // Debugger Notify Halt XFX-form (dnh DUI,DUIS)
		[5]*argField{ap_ImmUnsigned_6_10, ap_ImmUnsigned_11_20}},
	{MSGSND, 0xfc0007fe, 0x7c00019c, 0x3ff0001, // Message Send X-form (msgsnd RB)
		[5]*argField{ap_Reg_16_20}},
	{MSGCLR, 0xfc0007fe, 0x7c0001dc, 0x3ff0001, // Message Clear X-form (msgclr RB)
		[5]*argField{ap_Reg_16_20}},
	{DCI, 0xfc0007fe, 0x7c00038c, 0x21ff801, // Data Cache Invalidate X-form (dci CT)
		[5]*argField{ap_ImmUnsigned_7_10}},
	{ICI, 0xfc0007fe, 0x7c00078c, 0x21ff801, // Instruction Cache Invalidate X-form (ici CT)
		[5]*argField{ap_ImmUnsigned_7_10}},
	{DCREAD, 0xfc0007fe, 0x7c0003cc, 0x1, // Data Cache Read X-form (dcread RT,RA,RB)
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ICREAD, 0xfc0007fe, 0x7c0007cc, 0x3e00001, // Instruction Cache Read X-form (icread RA,RB)
		[5]*argField{ap_Reg_11_15, ap_Reg_16_20}},
	{MFPMR, 0xfc0007fe, 0x7c00029c, 0x1, // Move From Performance Monitor Register XFX-form (mfpmr RT,PMRN)
		[5]*argField{ap_Reg_6_10, ap_SpReg_11_20}},
	{MTPMR, 0xfc0007fe, 0x7c00039c, 0x1, // Move To Performance Monitor Register XFX-form (mtpmr PMRN,RS)
		[5]*argField{ap_SpReg_11_20, ap_Reg_6_10}},
}
