//
//  VersionBumper.swift
//  
//
//  Created by Kamaal M Farah on 10/10/2021.
//

import Foundation
import ArgumentParser

@main
struct VersionBumper: ParsableCommand {
    @Option(name: .shortAndLong, help: "Build number")
    var buildNumber: Int?
    @Option(name: .shortAndLong, help: "version")
    var version: String?
    @Option(name: .shortAndLong, help: "Info.plist path")
    var infoPlist: String

    enum Errors: Error {
        case invalidInfoPlistURLError
        case contentNotFound
        case buildTagNotFound
        case versionTagNotFound
    }

    func run() throws {
        try timeFunction {
            guard let infoPlistURL = URL(string: infoPlist) else { throw Errors.invalidInfoPlistURLError }
            guard let infoPlistDictionary = NSMutableDictionary(contentsOfFile: infoPlistURL.path) else {
                throw Errors.contentNotFound
            }

            if let buildNumber = self.buildNumber {
                let buildNumberKey = "CFBundleVersion"
                guard infoPlistDictionary[buildNumberKey] != nil else {
                    throw Errors.buildTagNotFound
                }
                infoPlistDictionary.setValue("\(buildNumber)", forKey: buildNumberKey)
            }

            if let version = self.version {
                let versionKey = "CFBundleShortVersionString"
                guard infoPlistDictionary[versionKey] != nil else {
                    throw Errors.versionTagNotFound
                }
                infoPlistDictionary.setValue(version, forKey: versionKey)
            }

            infoPlistDictionary.write(toFile: infoPlistURL.path, atomically: false)
        }
    }

    private func timeFunction(completion: (() throws -> Void)) rethrows {
        let start = CFAbsoluteTimeGetCurrent()

        try completion()

        let diff = CFAbsoluteTimeGetCurrent() - start
        print("Took \(diff) seconds ✨")
    }
}
