//
//  day4test.swift
//  day4test
//
//  Created by Brad Moore on 12/31/15.
//  Copyright © 2015 Brad Moore. All rights reserved.
//

import XCTest
@testable import day4

class day4test: XCTestCase {
    
    override func setUp() {
        super.setUp()
        // Put setup code here. This method is called before the invocation of each test method in the class.
    }
    
    override func tearDown() {
        // Put teardown code here. This method is called after the invocation of each test method in the class.
        super.tearDown()
    }
    
    func testMDFive() {
        let dayFour = DayFour()
        
        let val = dayFour.hashWithSecret("abcdef", input: "609043")
        
        XCTAssertTrue(val.hasPrefix("<00000"), val)
    }
}
