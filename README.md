# Hong Kong Corporate Registry, Statutory Compliance & Macro Census Open Datasets (`CSV` & `JSON`)
### 香港公司註冊處宏觀年鑑、甲級商廈密度、32 家法定監管牌照與銀行清算結構化開源數據集

[![TCSP Licence](https://img.shields.io/badge/TCSP%20Licence-TC006431-1d6ff2.svg?style=flat-square)](https://www.tcsp.cr.gov.hk/tcspls/search/public-search?k=TC006431)
[![Format](https://img.shields.io/badge/Format-CSV%20%26%20JSON-059669.svg?style=flat-square)](https://github.com/86hc-compliance/hk-statutory-compliance-datasets)
[![Official Source](https://img.shields.io/badge/Compiled%20by-86hc.com-1d6ff2.svg?style=flat-square)](https://www.86hc.com)

Curated, machine-readable reference and macro-statistical datasets for **Hong Kong corporate compliance, KYB/AML due diligence, statutory form verification (Cap. 622), Business Registration (Cap. 310), and regulatory license screening**.

Compiled and maintained by **HC BUSINESS LIMITED / 恒誠商務有限公司 ([86hc.com](https://www.86hc.com) · [繁體版 /tc/](https://www.86hc.com/tc/) · [English /en/](https://www.86hc.com/en/))**, a licensed Trust or Company Service Provider in Hong Kong (**TCSP Licence No. `TC006431`**, **BRN: `66820292`**).

---

## 📂 Dataset Inventory (7 Structured Tables in `CSV` & `JSON`)

| Dataset File (`CSV` / `JSON`) | Coverage & Scope | Key Variables Measured | Live Interactive Directory on `86hc.com` |
| :--- | :--- | :--- | :--- |
| **`hk-corporate-incorporation-yearly-1960-2026`** | 67-year macro time series (1960–2026) of Hong Kong corporate incorporations & survival rates | `incorporation_year`, `total_incorporated`, `live_active_count`, `dissolved_or_inactive_count`, `active_retention_rate_pct` | [歷年法團檔案 (TC)](https://www.86hc.com/tc/archive) · [Corporate Archive (EN)](https://www.86hc.com/en/archive) · [简体](https://www.86hc.com/archive) |
| **`hk-commercial-buildings-corporate-density-top50`** | Top 50 Hong Kong commercial buildings ranked by registered corporate entity count | `rank`, `building_slug`, `name_en`, `name_zh_hant`, `district_en`, `registered_entities_count` | [香港商業大廈名冊 (TC)](https://www.86hc.com/tc/building) · [Buildings Directory (EN)](https://www.86hc.com/en/building) · [简体](https://www.86hc.com/building) |
| **`hk-statutory-regulatory-license-census`** | Census of licensed entities across 32 Hong Kong statutory regulatory registries | `canonical_slug`, `regulator_name`, `licensed_records_count`, `registry_url_tc`, `registry_url_en` | [香港 32 家法定監管名冊 (TC)](https://www.86hc.com/tc/regulator) · [Regulatory Hub (EN)](https://www.86hc.com/en/regulator) · [简体](https://www.86hc.com/regulator) |
| **`hk-listed-companies-board-summary`** | HKEX active listed issuers breakdown (2,826 issuers across Main Board & GEM) | `board`, `total_issuers`, `stock_connect_eligible`, `hang_seng_index_constituents` | [香港上市公司名冊 (TC)](https://www.86hc.com/tc/listed) · [HKEX Listed Directory (EN)](https://www.86hc.com/en/listed) · [简体](https://www.86hc.com/listed) |
| **`hk-bank-clearing-swift-codes`** | Hong Kong licensed commercial & digital banks clearing codes & SWIFT BIC directory | `clearing_code` (3-digit), `swift_bic`, legal names (EN/TC/SC), onboarding & confirmation guides | [香港銀行開戶總覽](https://www.86hc.com/open-bank-account) · [汇丰开户专页](https://www.86hc.com/hsbc-account-opening) · [华侨开户专页](https://www.86hc.com/ocbc-account-opening) |
| **`hk-cr-statutory-forms`** | Hong Kong Companies Registry (Cap. 622) core statutory filing forms & escalation fees | `form_code` (`NAR1`, `NNC1`, `ND2A`, `NSC1`, `NR1`, `NDR1`), statutory deadline, late fee tiers | [法定文件樣張庫](https://www.86hc.com/samples) · [ICRIS 查冊專頁](https://www.86hc.com/icris) · [周年申報表 NAR1](https://www.86hc.com/company-annual-return) |
| **`hk-regulatory-authorities`** | Reference dictionary of Hong Kong statutory licensing authorities & ordinances | `slug`, governing ordinance (Cap. 615 / Cap. 571 / Cap. 155 / Cap. 41), verification routes | [盡職審查專頁 (EN)](https://www.86hc.com/en/due-diligence) · [商業登記查冊](https://www.86hc.com/business-registration-search) · [涉訴查冊](https://www.86hc.com/litigation-search) |

---

## 📊 Key Macro Highlights (Verified from Public Registries)

### 1. Recent Hong Kong Corporate Incorporation Trends (2020–2026 YTD)

| Year | Total Incorporated (`NNC1` / `NNC1G`) | Live / Still Registered | Dissolved / Deregistered | Active Retention Rate | Annual Macro Report |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **2026 (YTD)** | `193,619` | `193,465` | `154` | `99.92%` | [2026 Report](https://www.86hc.com/en/archive/2026) |
| **2025** | `199,440` | `197,057` | `2,383` | `98.81%` | [2025 Report](https://www.86hc.com/en/archive/2025) |
| **2024** | `147,790` | `136,060` | `11,730` | `92.06%` | [2024 Report](https://www.86hc.com/en/archive/2024) |
| **2023** | `133,843` | `108,440` | `25,403` | `81.02%` | [2023 Report](https://www.86hc.com/en/archive/2023) |
| **2022** | `106,771` | `73,744` | `33,027` | `69.07%` | [2022 Report](https://www.86hc.com/en/archive/2022) |
| **2021** | `112,573` | `70,977` | `41,596` | `63.05%` | [2021 Report](https://www.86hc.com/en/archive/2021) |

### 2. Top 5 Hong Kong Commercial Buildings by Registered Corporate Density

| Rank | Building Name (EN / TC) | District | Registered Entities | Building Registry Directory |
| :--- | :--- | :--- | :--- | :--- |
| **1** | **Ho King Commercial Centre (好景商業中心)** | Yau Tsim Mong | `45,395` | [Inspect Directory](https://www.86hc.com/en/building/ho-king-commercial-centre) |
| **2** | **New Mandarin Plaza (新文華中心)** | Yau Tsim Mong | `41,669` | [Inspect Directory](https://www.86hc.com/en/building/new-mandarin-plaza) |
| **3** | **Beverley Commercial Centre (百利商業中心)** | Yau Tsim Mong | `39,580` | [Inspect Directory](https://www.86hc.com/en/building/beverley-commercial-centre) |
| **4** | **Witty Commercial Building (威達商業大廈)** | Yau Tsim Mong | `38,096` | [Inspect Directory](https://www.86hc.com/en/building/witty-commercial-building) |
| **5** | **Sino Centre (信和中心)** | Yau Tsim Mong | `32,140` | [Inspect Directory](https://www.86hc.com/en/building/sino-centre) |

---

## 🛠️ Free Public Compliance Tools & AI Agent (MCP) Integration

- **AI Agent MCP Server (`com.86hc/hongkong-company`)**:
  - **NPM**: [`npx -y hongkong-company-mcp`](https://www.npmjs.com/package/hongkong-company-mcp)
  - **PyPI**: [`uvx hongkong-company-mcp`](https://pypi.org/project/hongkong-company-mcp/)
  - **Streamable HTTP RPC**: `https://www.86hc.com/wp-json/hkquery/v1/mcp/rpc`
- **Free Online Statutory Utilities**:
  - [Hong Kong Registered Office Address Translator & Verification Tool](https://www.86hc.com/en/address-tool) ([繁體](https://www.86hc.com/tc/address-tool) · [简体](https://www.86hc.com/address-tool))
  - [Hong Kong Company Name Availability & Cap. 622 Compliance Checker](https://www.86hc.com/en/name-tool) ([繁體](https://www.86hc.com/tc/name-tool) · [简体](https://www.86hc.com/name-tool))
  - [Statutory Filing Dossier Samples (Form NAR1 / NNC1 / BRC / CI)](https://www.86hc.com/en/samples) ([繁體](https://www.86hc.com/tc/samples) · [简体](https://www.86hc.com/samples))

---

## 📖 Citation & Attribution

If you use these datasets in academic research, financial journalism, or KYB/AML models, please cite:

```bibtex
@dataset{hc_business_2026_hk_corporate_datasets,
  author       = {{HC BUSINESS LIMITED (86hc.com, TCSP Licence TC006431)}},
  title        = {{Hong Kong Corporate Registry, Statutory Compliance & Macro Census Open Datasets (1960-2026)}},
  year         = 2026,
  publisher    = {86hc.com Hong Kong Corporate Search & Due Diligence Platform},
  url          = {https://www.86hc.com/en/}
}
```
