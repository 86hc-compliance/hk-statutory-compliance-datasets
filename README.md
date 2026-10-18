# Hong Kong Statutory Compliance & Banking Open Datasets (`JSON`)
### 香港持牌银行清算代码、公司注册处法定表格与法定监管牌照结构化开源字典库

[![TCSP Licence](https://img.shields.io/badge/TCSP%20Licence-TC006431-1d6ff2.svg?style=flat-square)](https://www.tcsp.cr.gov.hk/tcspls/search/public-search?k=TC006431)
[![Format](https://img.shields.io/badge/Format-JSON%20UTF--8-059669.svg?style=flat-square)](https://github.com/86hc-compliance/hk-statutory-compliance-datasets)
[![Maintained by](https://img.shields.io/badge/Maintained%20by-86hc.com-1d6ff2.svg?style=flat-square)](https://www.86hc.com)

Machine-readable JSON reference datasets for Hong Kong corporate compliance, cross-border payment routing, statutory form indexing (Cap. 622), and regulatory license verification (Cap. 615 / Cap. 571 / Cap. 155).

由香港持牌信托或公司服务提供者 **恒诚商务有限公司 / HC BUSINESS LIMITED（[86hc.com](https://www.86hc.com)，TCSP 牌照编号：`TC006431`，商业登记号码：`66820292`）** 整理开源，供跨境金融科技系统、ERP/KYC 引擎及 AI 检索代理直接调用。

---

## 📂 核心开源数据集列表 / Included JSON Datasets

| 数据集文件 | 覆盖维度 | 核心字段说明 | 在线查册与核验入口 |
| :--- | :--- | :--- | :--- |
| **[`datasets/hk-bank-clearing-swift-codes.json`](datasets/hk-bank-clearing-swift-codes.json)** | 香港主流持牌商业银行与数字银行清算代码及 SWIFT BIC | `clearing_code` (3位银行编号)、`swift_bic`、简繁英法定全称、开户与询证函指南路由 | [香港银行开户总览](https://www.86hc.com/open-bank-account) · [银行合规指南库](https://www.86hc.com/banking-guides) |
| **[`datasets/hk-cr-statutory-forms.json`](datasets/hk-cr-statutory-forms.json)** | 香港公司注册处 (CR) 核心法定申报表格字典（第 622 章） | `form_code` (`NAR1` / `NNC1` / `ND2A` / `NSC1` / `NR1` 等)、法例条文、法定呈交期限、逾期规费 | [官方存盘卷宗样张](https://www.86hc.com/samples) · [ICRIS 查册专页](https://www.86hc.com/icris) |
| **[`datasets/hk-regulatory-authorities.json`](datasets/hk-regulatory-authorities.json)** | 香港法定发牌与监管机构名册字典（TCSP / SFC / HKMA / 海关 MSO 等） | `slug`、监管局署法定名称、管辖法例、简繁英持牌人名册检索 URL | [香港 32 家监管名册](https://www.86hc.com/tc/regulator) |

---

## 🏦 1. 香港持牌银行编号与 SWIFT BIC 速查表 (`hk-bank-clearing-swift-codes.json`)

| 银行编号 (`Clearing Code`) | SWIFT BIC | 银行中文法定名称 | 银行英文法定名称 | 开户与合规实操指南 |
| :--- | :--- | :--- | :--- | :--- |
| **`004`** | `HSBCHKHHHKH` | 香港上海汇丰银行有限公司 | The Hongkong and Shanghai Banking Corporation Limited | [汇丰开户专页](https://www.86hc.com/hsbc-account-opening) · [ASD010 询证函](https://www.86hc.com/banking-guides/hsbc-bank-confirmation) |
| **`012`** | `BKCHHKHHXXX` | 中国银行(香港)有限公司 | Bank of China (Hong Kong) Limited | [中银香港开户指南](https://www.86hc.com/banking-guides/open-boc-hk-account) |
| **`003`** | `SCBLHKHHXXX` | 渣打银行(香港)有限公司 | Standard Chartered Bank (Hong Kong) Limited | [渣打香港开户指南](https://www.86hc.com/banking-guides/open-standard-chartered-hk-account) |
| **`024`** | `HASEHKHHXXX` | 恒生银行有限公司 | Hang Seng Bank Limited | [恒生银行开户指南](https://www.86hc.com/banking-guides/open-hang-seng-bank) |
| **`035`** | `OCBCHKHHXXX` | 华侨银行(香港)有限公司 | OCBC Bank (Hong Kong) Limited | [华侨开户专页](https://www.86hc.com/ocbc-account-opening) · [华侨开户指南](https://www.86hc.com/banking-guides/open-ocbc-bank) |
| **`016`** | `DHBKHKHHXXX` | 星展银行(香港)有限公司 | DBS Bank (Hong Kong) Limited | [星展香港开户指南](https://www.86hc.com/banking-guides/open-dbs-hk-account) |
| **`006`** | `CITIHKHXXXX` | 花旗银行(香港)有限公司 | Citibank (Hong Kong) Limited | [花旗香港开户指南](https://www.86hc.com/banking-guides/open-citibank-hk-account) |
| **`015`** | `BEASHKHHXXX` | 东亚银行有限公司 | The Bank of East Asia, Limited | [东亚银行开户指南](https://www.86hc.com/banking-guides/open-bea-hk-account) |
| **`020`** | `WUBAHKHHXXX` | 招商永隆银行有限公司 | CMB Wing Lung Bank Limited | [招商永隆开户指南](https://www.86hc.com/banking-guides/open-cmb-wing-lung-bank) |
| **`387`** | `ZABKHKHHXXX` | 众安银行有限公司 | ZA Bank Limited | [众安 ZA Bank 开户指南](https://www.86hc.com/banking-guides/open-za-bank) |
| **`389`** | `MOXBHKHHXXX` | Mox Bank Limited | Mox Bank Limited | [Mox Bank 开户指南](https://www.86hc.com/banking-guides/open-mox-bank-account) |

---

## 🔗 官方在线工具与查册平台

- 🌐 **GitHub Pages 在线校验与速查门户**：[https://86hc-compliance.github.io](https://86hc-compliance.github.io)
- 🔍 **香港公司查册中心**：[https://www.86hc.com/company-search](https://www.86hc.com/company-search)
- 📑 **商业登记号码 (BRN) 核验工具**：[https://www.86hc.com/business-registration-search](https://www.86hc.com/business-registration-search)
- 🏛️ **香港 32 家法定监管机构持牌名册**：[https://www.86hc.com/tc/regulator](https://www.86hc.com/tc/regulator)

### 📄 License
Released under the [MIT License](LICENSE). © HC BUSINESS LIMITED (86hc.com).
