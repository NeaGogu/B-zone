// External dependencies
import React, { useState } from 'react';
import { MapContainer, TileLayer, Marker, Popup, useMapEvents } from 'react-leaflet';
import { LayersControl, LayerGroup } from 'react-leaflet';
import 'leaflet/dist/leaflet.css';
import 'leaflet-defaulticon-compatibility/dist/leaflet-defaulticon-compatibility.css';
import "leaflet-defaulticon-compatibility";

// Components
import Heatmap from './heatmapComponent';
import PolygonVis from './polygonComponents'

/**
 * Shows address, zipcode and coordinates when clicking on the map.
 * @return {JSX.Element|null} - Returns a JSX element containing a map marker with a popup, or null if the position is null.
 */
function LocationMarker() {
    const [position, setPosition] = useState(null);
    const [address, setAddress] = useState(null);
    const [zipcode, setZipcode] = useState(null);

    const map = useMapEvents({
        click(e) {
            const { lat, lng } = e.latlng;
            setPosition(e.latlng);
            const API_KEY = 'pk.eyJ1IjoidGFuaWFnb2lhMTEiLCJhIjoiY2xleTRrYm02MDlmMTN4bzVsZTR4cWp4OCJ9.hmT59q-Q1IcEjC6mdY2R9w';
            const API_URL = `https://api.mapbox.com/geocoding/v5/mapbox.places/${lng},${lat}.json?access_token=${API_KEY}&types=postcode,address`;
            fetch(API_URL)
                .then(response => response.json())
                .then(data => {
                    const features = data.features;
                    if (features.length > 0) {
                        const address = features[0].place_name;
                        const zipcodes = features.filter(feature => feature.place_type[0] === 'postcode');
                        const zipcode = zipcodes.length > 0 ? zipcodes[0].text : null;
                        setAddress(address);
                        setZipcode(zipcode);
                    } else {
                        setAddress(null);
                        setZipcode(null);
                    }
                });
            map.flyTo(e.latlng, map.getZoom());
        },
    });

    return position === null ? null : (
        <Marker position={position}>
            <Popup>
                <div>
                    <div>Latitude: {position.lat.toFixed(4)}</div>
                    <div>Longitude: {position.lng.toFixed(4)}</div>
                    {address && <div>Address: {address}</div>}
                    {zipcode && <div>Zipcode: {zipcode}</div>}
                </div>
            </Popup>
        </Marker>
    );
}


// Main function to hold the map, location marker and the layers
function MapComponent(props) {
    const { value, intensity } = props;

    return (
        <MapContainer center={[52, 7]} zoom={7} scrollWheelZoom={true} style={{ height: 500, flex: "1" }}>
            <TileLayer attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
                       url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" />

            <LayersControl position='topright'>
                <LayersControl.Overlay name='Heat map'>
                    <LayerGroup>
                        <Heatmap value={value} intensity={intensity} />
                    </LayerGroup>
                </LayersControl.Overlay>

                <LayersControl.Overlay name='Zones'>
                    <LayerGroup>
                        <PolygonVis />
                    </LayerGroup>
                </LayersControl.Overlay>
            </LayersControl>

            {/*PolygonVis*/}

            <LocationMarker />
        </MapContainer>
    );
}

export default MapComponent;