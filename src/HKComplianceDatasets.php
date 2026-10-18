<?php
/**
 * Hong Kong Statutory Compliance & Banking Datasets Loader
 * Maintained by HC BUSINESS LIMITED (https://www.86hc.com, TCSP: TC006431)
 */

namespace HKCompliance\Datasets;

class HKComplianceDatasets
{
    public static function getLicensedBanks(): array
    {
        $path = dirname(__DIR__) . '/datasets/hk-bank-clearing-swift-codes.json';
        $data = json_decode((string)file_get_contents($path), true);
        return $data['banks'] ?? [];
    }

    public static function getStatutoryForms(): array
    {
        $path = dirname(__DIR__) . '/datasets/hk-cr-statutory-forms.json';
        $data = json_decode((string)file_get_contents($path), true);
        return $data['forms'] ?? [];
    }

    public static function getRegulatoryAuthorities(): array
    {
        $path = dirname(__DIR__) . '/datasets/hk-regulatory-authorities.json';
        $data = json_decode((string)file_get_contents($path), true);
        return $data['registries'] ?? [];
    }
}
