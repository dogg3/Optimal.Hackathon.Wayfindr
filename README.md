# OptimALL - Wayfindr’s Public Transport Hackathon 2.0 https://www.wayfindr.net/

OptimALL is a multisensor data fusion tool that uses currently available monitoring systems in the Underground to help TfL staff manage the Tube network. We combine data from multiple sensors in TfL Underground stations in a web application that staff can use to intuitively and remotely monitor how busy a station is, and make key decisions on station management. 

The Problem

Wayfindr mostly relies on beacons to communicate directions to its users. This system can be subject to multiple issues, such as low battery level, inaccuracies, signal failure, and even delivering the wrong instructions.
This can result in incomplete journeys and users frustration; and with no back up system, this also adds extra responsibilities for TfL staff, who then have to take care of the matter themselves as soon as possible.

The Solution

In that context, OptimALL offers an analytical tool that uses other available monitoring systems - such as CCTV and fare gates - to provide additional information about users' journey. The extra source of data collected can feed into the Wayfindr app as an alternative way to keep track of users, and make sure they are on the right path. It can also provide meaningful information for TfL staff - i.e. congestion warnings, route optimisation, etc. - as well as improving key performance metrics - i.e. EJT (Estimated Journey Time) and LCH (Lost Customer Hours) - and reduce the need to intervene in case of a beacon failure.

The Technology

OptimALL can collect data from multiple existing sources from TfL. However for the sake of this hackathon, we will focus on data provided by CCTVs. We use facial recognition to keep track of users across the underground tunnels. These data can then be cross-checked against Wayfindr's insure the validity and accuracy of the navigation system. In addition, TfL can use those data to make informed decisions for the optimisation of customers’ journey.

### Prerequisites

To implement our facial recognition system, the following prerequisites must be installed.

brew install cmake
pip3 install dlib
pip3 install opencv-python
python3 image_recognition_script.py

### CCTV Facial recognition

Facial recognition is implemented by modifying the OpenCV library and Dlib.

### Web application

Our web application is built using Java EE, running within a GlassFish web server.

Congestion Maps
We demonstrate our method of optimal path selection through the network of walkways in a station. This is based on the aggregate data of multiple Wayfindr beacons, CCTV cameras, ticket barrier data and infrared sensors. 

Sensor Performance
Each sensor type should be correlated with time, thus, if not, that indicates the presence of a fault with one or more sensors. This is a self-regulating system that provides reliable 

Station Analytics
This tab shows the overall station performance. The aggregated data from all passengers is displayed as total footfall per hour, and average time spent in the station, both of which implies the business of the station.

## Authors

* **Andrea**
* **Douglas**
* **Gavin**
* **Matthew**
* **Pierre-Antoine**
* **Osman**

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

Thanks Wayfindr's Public Transport Hackathon 2.0 for the event!
