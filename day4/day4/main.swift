//
//  main.swift
//  day4
//
//  Created by Brad Moore on 12/31/15.
//  Copyright © 2015 Brad Moore. All rights reserved.
//

import Foundation

class DayFour {
    func md5(string string: String) -> NSData {
        let digest = NSMutableData(length: Int(CC_MD5_DIGEST_LENGTH))!
        if let data :NSData = string.dataUsingEncoding(NSUTF8StringEncoding) {
            CC_MD5(data.bytes, CC_LONG(data.length),
                UnsafeMutablePointer<UInt8>(digest.mutableBytes))
        }
        return digest
    }
    
    func main() {
        let secret = "iwrupvqb"
        
        var answer: String = ""
        
        var i: Int = 0
        while (true) {
            
            let hash = hashWithSecret(secret, input: String(i))
            if (hash.hasPrefix("<000000")) {
                answer = hash
                break
            }

            i += 1
        }
        
        print("Input \(i) makes hash \(answer)")
    }
    
    func hashWithSecret(secret: String, input: String) -> String {
        let mdFiveData = md5(string: secret + input)
        
        return String(mdFiveData)
    }
}

var dayFour = DayFour()
dayFour.main()