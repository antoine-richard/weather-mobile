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

                int resId = getResources().getIdentifier("city"+(i+1)+"name", "id", getContext().getPackageName());
                ((TextView) rootView.findViewById(resId)).setText(city.getString("name"));

                resId = getResources().getIdentifier("city"+(i+1)+"description", "id", getContext().getPackageName());
                ((TextView) rootView.findViewById(resId)).setText(city.getString("desc"));

                resId = getResources().getIdentifier("city"+(i+1)+"temperature", "id", getContext().getPackageName());
                ((TextView) rootView.findViewById(resId)).setText(city.getString("temp"));

            }
        } catch (Exception e) {
            e.printStackTrace();
        }

        return rootView;
    }
}