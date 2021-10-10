// swift-tools-version:5.4
// The swift-tools-version declares the minimum version of Swift required to build this package.

import PackageDescription

let package = Package(
    name: "VersionBumper",
    products: [
        .library(name: "VersionBumper", targets: [
            "version-bumper",
        ]),
    ],
    dependencies: [
        .package(url: "https://github.com/apple/swift-argument-parser", "1.0.0"..<"2.0.0"),
    ],
    targets: [
        .executableTarget(name: "version-bumper", dependencies: [
            .product(name: "ArgumentParser", package: "swift-argument-parser"),
        ]),
        .testTarget(name: "VersionBumperTests", dependencies: [
            "version-bumper"
        ]),
    ]
)
