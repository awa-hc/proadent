<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use app\Models\Appointment;

class AppointmentController extends Controller
{
    public $nro;

    public function index()
    {
        $appointments = Appointment::all();
        return response()->json($appointments, 200);
    }

    public function store(Request $request)
    {
        $validator = Validator::make($request->all(), [
            'user_id' => 'required|exists:users,id',
            'type' => 'required|integer',
            'date' => 'required|dateTime',
            'reason' => 'required|string|max:255',
        ]);

        if ($validator->fails())
        {
            return response()->json(['errors' => $validator->errors()], 422);
        }

        $codigo = $this->generarNroCita($request);

        $appointment = Appointment::create([
            'user_id' => $request->user_id,
            'nro' => $this->nro,
            'codigo' => $codigo,
            'type' => $request->type,
            'date' => $request->date,
            'reason' => $request->reason,
            'status' => 1,
        ]);

        return response()->json($appointment, 201);
    }

    public function show(Appointment $appointment)
    {
        return response()->json($appointment, 200);
    }

    public function update(Request $request, Appointment $appointment)
    {
        $request->validate([
            'user_id' => 'required|exists:users,id',
            'type' => 'required|integer',
            'date' => 'required|dateTime',
            'reason' => 'required|string|max:255',
        ]);

        $appointment->update($request->all());

        return response()->json($appointment, 200);
    }

    public function destroy(Appointment $appointment)
    {
        $appointment->delete();

        return response()->json(null, 204);
    }

    private function generarNroCita(Request $request)
    {
        $nuevoNroCitaMedica = "";
        $date = Carbon::parse($request->date);
        $ultimoNro = Appointment::whereYear('fecha', $date->year)
            ->orderBy('nro', 'DESC')
            ->first();
        if ($ultimoNro != null)
        {
            $this->nro = intVal($ultimoNro->nro) + 1;
        }
        else
        {
            $this->nro = 1;
        }

        $nuevoNroCitaMedica = $nuevoNroCitaMedica . 'CM';

        $nuevoNroCitaMedica .= '-' . $date->year . '-' . str_pad($this->nro, 6, '0', STR_PAD_LEFT);

        return $nuevoNroCitaMedica;
    }

    public function realizarCita(Request $request)
    {
        $cita = Appointment::find($request->id);
        //if ($this->user->can('caja-chica-admin.aprobar') || $cajaChica->estado != CajaChica::PENDIENTE)
        //{
        $cita->update([
            'estado' => Appointment::REALIZADA,
            'user_auth_id' => Auth::user()->id,
        ]);
        return response()->json($cita, 200);
        //}
        //else
        //{
        //    $this->emit('msg-error', 'Acción no autorizada!');
        //}
    }
}
