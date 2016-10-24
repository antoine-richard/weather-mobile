package com.github.antoinerichard.weathermobile.androidapp;

import android.os.Bundle;
import android.support.v4.app.Fragment;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;

import org.json.JSONArray;
import org.json.JSONObject;

import go.weather.Weather;

public class HomeFragment extends Fragment {

    @Override
    public View onCreateView(LayoutInflater inflater, ViewGroup container,
                             Bundle savedInstanceState) {
        View rootView = inflater.inflate(R.layout.fragment_home, container, false);

        try {
            byte[] data = Weather.fetchDefaultCities();
            String json = new String(data);
            JSONArray cityList = new JSONArray(json);
            for (int i = 0; i < cityList.length(); i++) {
                JSONObject city = cityList.getJSONObject(i);
//                    labels.get(i).setText(
                TextView cityView = (TextView) rootView.findViewById(R.id.city);
                cityView.setText(city.getString("name"));
                TextView descriptionView = (TextView) rootView.findViewById(R.id.description);
                descriptionView.setText(city.getString("desc"));
                TextView temperatureView = (TextView) rootView.findViewById(R.id.temperature);
                temperatureView.setText(city.getString("temp"));
            }
        } catch (Exception e) {
            e.printStackTrace();
        }

        return rootView;
    }
}