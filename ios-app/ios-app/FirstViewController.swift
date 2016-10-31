//
//  FirstViewController.swift
//  ios-app
//
//  Created by Antoine Richard on 25/10/16.
//  Copyright © 2016 Antoine Richard. All rights reserved.
//

import UIKit
import Weather

class FirstViewController: UIViewController {

    @IBOutlet weak var city1Name: UILabel!
    @IBOutlet weak var city1Temperature: UILabel!
    @IBOutlet weak var city1Description: UILabel!
    
    @IBOutlet weak var city2Name: UILabel!
    @IBOutlet weak var city2Temperature: UILabel!
    @IBOutlet weak var city2Description: UILabel!
    
    @IBOutlet weak var city3Name: UILabel!
    @IBOutlet weak var city3Temperature: UILabel!
    @IBOutlet weak var city3Description: UILabel!
    
    var labels:[[String: UILabel]]=[[:]]
    
    @IBOutlet weak var errorMessage: UILabel!
    
    @IBOutlet weak var weatherStack: UIStackView!
    @IBOutlet weak var errorStack: UIStackView!
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        labels = [[
                "name": city1Name,
                "temp": city1Temperature,
                "desc": city1Description
            ], [
                "name": city2Name,
                "temp": city2Temperature,
                "desc": city2Description
            ], [
                "name": city3Name,
                "temp": city3Temperature,
                "desc": city3Description
            ]
        ]

        fetchFeaturedWeather()
    }

    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
        // Dispose of any resources that can be recreated.
    }

    func fetchFeaturedWeather() {
        var err: NSError?
        var weatherData: NSData?
        _ = Weather.GoWeatherFetchDefaultCities(&weatherData, &err)
        
        if err != nil {
            errorMessage.text = err!.localizedDescription
            errorStack.isHidden = false
            weatherStack.isHidden = true
        } else {
            let json = try? JSONSerialization.jsonObject(with: weatherData! as Data, options: [])
            if let cities = json as? [[String: String]] {
                for (index, city) in cities.enumerated() {
                    labels[index]["name"]!.text = city["name"]
                    labels[index]["temp"]!.text = city["temp"]
                    labels[index]["desc"]!.text = city["desc"]
                }
                errorStack.isHidden = true
                weatherStack.isHidden = false
            } else {
                errorMessage.text = "Unable to deserialize weather :("
                errorStack.isHidden = false
                weatherStack.isHidden = true
            }
        }
    }
    
    @IBAction func retry() {
        errorStack.isHidden = true
        fetchFeaturedWeather()
    }

}

