<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class VerificationTokenController extends Controller
{
    public function generateToken($user_id)
    {
        $token = str_random(60);

        VerificationToken::create([
            'user_id' => $user_id,
            'token' => $token,
        ]);

        return response()->json(['token' => $token], 200);
    }
}
