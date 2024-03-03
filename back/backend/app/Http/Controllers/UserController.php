<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Models\User;
use Illuminate\Validation\Rule;
use Illuminate\Support\Facades\Validator;
use Illuminate\Support\Facades\Http;

class UserController extends Controller
{
    /**
     *
     * @param  \Illuminate\Http\Request  $request
     * @return \Illuminate\Http\Response
     */
    public function register(Request $request)
    {
        $validator = Validator::make($request->all(), [
            'name' => 'required|string|max:255',
            'email' => 'sometimes|required_without:phone|string|email|max:255|unique:users',
            'phone' => 'sometimes|required_without:email|string|regex:/^[0-9]+$/|max:11',
            'password' => 'required|string|min:8',
            'password_confirmation' => 'required|string|same:password',
        ]);

        if ($validator->fails()) {
            return response()->json(['errors' => $validator->errors()], 422);
        }


        $phone = $request->input('phone');
        $email = $request->input('email');

        if ($phone) {
            $phone = '+591' + $phone;
        }

        $user = User::create([
            'name' => $request->name,
            'birthday' => $request->birthday,
            'ci' => $request->ci,
            'phone' => $phone,
            'email' => $request->email,
            'password' => bcrypt($request->password),

        ]);

        if ($email) {
            $this->sendEmail($email);
        } elseif ($phone) {
            $this->sendPhone($phone);
        }

        return response()->json(['user' => $user], 201);
    }


    public function sendEmail(string $email)
    {
        $response = Http::withOptions([
            'verify' => false,
        ])->post(
                'http://localhost:8080/email',
                [
                    'email' => $email,
                ]
            );
        if ($response->successful()) {
            return 'Email enviado correctamente';
        } else {
            return 'Error al enviar el correo';
        }
    }

    public function sendPhone(string $phone)
    {
        $response = Http::withOptions([
            'verify' => false,
        ])->post(
                'https://dc34sk6l-8080.brs.devtunnels.ms/phone',
                [
                    'phone' => $phone,
                ]
            );
        if ($response->successful()) {
            return 'Phone enviado correctamente';
        } else {
            return 'Error al enviar el correo';
        }
    }

    public function update(Request $request, User $user)
    {
        $validator = Validator::make($request->all(), [
            'name' => 'required|string|max:255',
            'email' => 'required|string|email|max:255|unique:users,email,' . $user->id,
            'phone' => 'required|string|max:11',
        ]);

        if ($validator->fails()) {
            return response()->json(['errors' => $validator->errors()], 422);
        }

        $user->update([
            'name' => $request->name,
            'birthday' => $request->birthday,
            'ci' => $request->ci,
            'phone' => $request->phone,
            'email' => $request->email,
        ]);

        return response()->json(['user' => $user], 200);
    }

    public function destroy(User $user)
    {
        $user->delete();

        return response()->json(null, 204);
    }
}
