<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

return new class extends Migration
{
    /**
     * Run the migrations.
     */
    public function up(): void
    {
        Schema::create('appointments', function (Blueprint $table)
        {
            $table->id();
            $table->foreignId('user_id')->constrained();
            $table->integer('nro');
            $table->string('codigo');
            $table->integer('type');
            $table->dateTime('date');
            $table->string('reason');
            $table->integer('status');
            $table->integer('user_auth_id');
            $table->timestamps();
        });
    }

    /**
     * Reverse the migrations.
     */
    public function down(): void
    {
        Schema::dropIfExists('appointments');
    }
};
