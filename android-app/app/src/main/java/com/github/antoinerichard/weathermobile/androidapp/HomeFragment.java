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
        TextView textView = (TextView) rootView.findViewById(R.id.section_label);

        try {
            byte[] data = Weather.fetchDefaultCities();
            String json = new String(data);
            JSONArray cityList = new JSONArray(json);
            for (int i = 0; i < cityList.length(); i++) {
                JSONObject city = cityList.getJSONObject(i);
//                    labels.get(i).setText(
                textView.setText(
                        city.getString("name") +
                                ", " +
                                city.getString("temp") +
                                ", " +
                                city.getString("desc"));
            }
        } catch (Exception e) {
            e.printStackTrace();
        }

        return rootView;
    }
}