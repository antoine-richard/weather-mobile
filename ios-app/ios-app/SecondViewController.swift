//
//  SecondViewController.swift
//  ios-app
//
//  Created by Antoine Richard on 25/10/16.
//  Copyright © 2016 Antoine Richard. All rights reserved.
//

import UIKit
import Weather

class SecondViewController: UIViewController, UISearchBarDelegate {
    
    @IBOutlet weak var searchBar: UISearchBar!
    
    @IBOutlet weak var cityName: UILabel!
    @IBOutlet weak var cityTemperature: UILabel!
    @IBOutlet weak var cityDescription: UILabel!
    
    @IBOutlet weak var errorMessage: UILabel!
    
    @IBOutlet weak var weatherStack: UIStackView!
    @IBOutlet weak var errorStack: UIStackView!
    
    override func viewDidLoad() {
        super.viewDidLoad()
        searchBar.delegate = self
        errorStack.isHidden = true
        weatherStack.isHidden = true
    }

    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
        // Dispose of any resources that can be recreated.
    }
    
    func searchBarSearchButtonClicked(_ searchBar: UISearchBar) {
        fetchCustomWeather(city: searchBar.text)
    }

    func fetchCustomWeather(city: String?) {
        var err: NSError?
        var weatherData: NSData?
        _ = Weather.GoWeatherFetchCustomCity(city, &weatherData, &err)
        
        if err != nil {
            errorMessage.text = err?.localizedDescription
            errorStack.isHidden = false
            weatherStack.isHidden = true
        } else {
            let json = try? JSONSerialization.jsonObject(with: weatherData! as Data, options: [])
            if let city = json as? [String: String] {
                cityName.text = city["name"]
                cityTemperature.text = city["temp"]
                cityDescription.text = city["desc"]
                errorStack.isHidden = true
                weatherStack.isHidden = false
            } else {
                errorMessage.text = "Unable to deserialize weather :("
                errorStack.isHidden = false
                weatherStack.isHidden = true
            }
        }
    }

}

