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
        $request->validate([
            'user_id' => 'required|exists:users,id',
            'procedure_id' => 'required|exists:procedures,id',
        ]);

        $price = Price::create($request->all());

        return response()->json($price, 201);
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
