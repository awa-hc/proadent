<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use app\Models\Price;

class PriceController extends Controller
{
    public function index()
    {
        $prices = Price::all();
        return response()->json($prices, 200);
    }

    public function store(Request $request)
    {
        $validator = Validator::make($request->all(), [
            'user_id' => 'required|exists:users,id',
            'procedure_id' => 'required|exists:procedures,id',
            'appointment_days' => 'required|integer|min:1',
            'total_price' => 'required|numeric|min:0',
        ]);

        if ($validator->fails())
        {
            return response()->json(['errors' => $validator->errors()], 422);
        }

        $price = Price::create([
            'user_id' => $request->user_id,
            'procedure_id' => $request->procedure_id,
            'appointment_days' => $request->appointment_days,
            'total_price' => $request->total_price,
        ]);

        return response()->json(['price' => $price], 201);
    }

    public function show(Price $price)
    {
        return response()->json($price, 200);
    }

    public function update(Request $request, Price $price)
    {
        $request->validate([
            'user_id' => 'required|exists:users,id',
            'procedure_id' => 'required|exists:procedures,id',
            'appointment_days' => 'required|integer|min:1',
            'total_price' => 'required|numeric|min:0',
        ]);

        $price->update($request->all());

        return response()->json($price, 200);
    }

    public function destroy(Price $price)
    {
        $price->delete();

        return response()->json(null, 204);
    }
}
