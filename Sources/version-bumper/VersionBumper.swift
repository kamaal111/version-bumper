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
    @Option(name: .shortAndLong, help: "Build number of app")
    var buildNumber: Int

    @Option(name: .shortAndLong, help: "Info.plist path")
    var infoPlist: String

    enum Errors: Error {
        case invalidInfoPlistURLError
        case contentNotFound
        case buildTagNotFound
    }

    func run() throws {
        try timeFunction {
            guard let infoPlistURL = URL(string: infoPlist) else { throw Errors.invalidInfoPlistURLError }
            guard let infoPlistDictionary = NSMutableDictionary(contentsOfFile: infoPlistURL.path) else {
                throw Errors.contentNotFound
            }
            let buildNumberKey = "CFBundleVersion"
            guard infoPlistDictionary[buildNumberKey] != nil else {
                throw Errors.buildTagNotFound
            }

            infoPlistDictionary.setValue("\(buildNumber)", forKey: buildNumberKey)

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
