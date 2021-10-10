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

    func run() throws {
        print("hello", buildNumber, infoPlist)
    }
}
