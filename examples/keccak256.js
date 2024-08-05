/**
 * ❯ for i in `cat services.txt`; do node examples/keccak256.js $i; done
 * 0x76dcbf2286d103a984031121f7e17e14579b0ef36f061320987117b5de138d04 :: cmp.services.ping.v1alpha.PingService
 * 0x100ed9b8ce3488f65518df64f1ea1721d91856f0530c7af7facc22ea36df4a26 :: cmp.services.partner.v1alpha.GetPartnerConfigurationService
 * 0x8319a0d4fe5612874de32e5e3fe9cca8ff46b0abab126db8b8e955fdb7892a2f :: cmp.services.transport.v1alpha.TransportSearchService
 * 0x3359014317cd4b2d832c8b6321a5e5282ec6dd6c3661928e17da32e80af7223e :: cmp.services.book.v1alpha.ValidationService
 * 0x72a5bff0ac1a58824d28736e28e174b56241c9affcbc139234eb580929f68e9d :: cmp.services.book.v1alpha.MintService
 * 0x168adba1a04cb5497c7fe14c241adbecf293ae463db5443f199705fb449ff70c :: cmp.services.activity.v1alpha.ActivityProductListService
 * 0x383a464bb5c6e678a0e2918de19c987bba292759a7d71f0ff558dbf0d0038718 :: cmp.services.activity.v1alpha.ActivitySearchService
 * 0xc34852694363aa7eb617f0b03ea03ea2e1fadd560c61c110714a71af0cda5ede :: cmp.services.info.v1alpha.CountryEntryRequirementsService
 * 0x3df535909c76cd0b18464f515206ac85b28f03ab7410716ba8d0abbd5fff3e0f :: cmp.services.network.v1alpha.GetNetworkFeeService
 * 0xc688cdc735bd2e69157527d9b7a9b7d2c385fb3975cd342a39f6039708886530 :: cmp.services.accommodation.v1alpha.AccommodationSearchService
 * 0x855eb588867b99d5b4373c597ab33e2ffc1d1254be10c5f9ce910f80d8521238 :: cmp.services.accommodation.v1alpha.AccommodationProductInfoService
 * 0x07880f944a3613912285ba367f1434abf8da9644fe9714052c81ee86ed1a4048 :: cmp.services.accommodation.v1alpha.AccommodationProductListService
 * 0xc41ac812d94725ea442e07fdbb2c06426f0f8221535bc932f1f45bf4105535ee :: cmp.services.seat_map.v1alpha.SeatMapService
 * 0x208b14833f44b92eac44315e1bb9b4a8e472c70270350be9c3f6abf1d1292eb7 :: cmp.services.seat_map.v1alpha.SeatMapAvailabilityService
 */
const ethers = require("ethers");

if (process.argv.length < 3) {
    console.error("Please provide a string to hash");
    process.exit(1);
}

const stringToHash = process.argv[2];

const hash = ethers.keccak256(ethers.toUtf8Bytes(stringToHash));
console.log(hash, "::", stringToHash);
