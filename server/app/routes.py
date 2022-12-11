from flask import request
import json
from app import app
from posapi import PosApi


@app.route('/')
def health():
    return "ebarimt service is running healthy!"


@app.route('/checkApi')
def checkApi():
    lib = request.args.get("lib")
    result = PosApi(lib).checkApi()
    return result


@app.route('/getInformation')
def getInformation():
    lib = request.args.get("lib")
    result = PosApi(lib).getInformation()
    return result


@app.route('/callFunction', methods=['POST'])
def callFunction():
    functionName = request.json['functionName']
    str_functionName = json.dumps(functionName)

    data = request.json['data']
    str_data = json.dumps(data)
    lib = request.args.get("lib")
    result = PosApi(lib).callFunction(
        str_functionName.encode('ascii'), str_data.encode('ascii'))

    return result


@app.route('/put', methods=['POST'])
def put():
    data = request.json['data']
    str_data = json.dumps(data)
    lib = request.args.get("lib")
    result = PosApi(lib).put(str_data.encode('ascii'))
    return result


@app.route('/returnBill', methods=['POST'])
def returnBill():
    data = request.json['data']
    str_data = json.dumps(data)
    lib = request.args.get("lib")
    result = PosApi(lib).returnBill(str_data.encode('ascii'))
    return result


@app.route('/sendData')
def sendData():
    lib = request.args.get("lib")
    result = PosApi(lib).sendData()
    return result
