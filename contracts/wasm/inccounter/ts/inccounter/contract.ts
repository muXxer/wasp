// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "wasmlib";
import * as sc from "./index";

export class CallIncrementCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncCallIncrement);
}

export class CallIncrementContext {
	state: sc.MutableIncCounterState = new sc.MutableIncCounterState();
}

export class CallIncrementRecurse5xCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncCallIncrementRecurse5x);
}

export class CallIncrementRecurse5xContext {
	state: sc.MutableIncCounterState = new sc.MutableIncCounterState();
}

export class EndlessLoopCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncEndlessLoop);
}

export class EndlessLoopContext {
	state: sc.MutableIncCounterState = new sc.MutableIncCounterState();
}

export class IncrementCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncIncrement);
}

export class IncrementContext {
	state: sc.MutableIncCounterState = new sc.MutableIncCounterState();
}

export class IncrementWithDelayCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncIncrementWithDelay);
	params: sc.MutableIncrementWithDelayParams = new sc.MutableIncrementWithDelayParams();
}

export class IncrementWithDelayContext {
	params: sc.ImmutableIncrementWithDelayParams = new sc.ImmutableIncrementWithDelayParams();
	state: sc.MutableIncCounterState = new sc.MutableIncCounterState();
}

export class InitCall {
	func: wasmlib.ScInitFunc = new wasmlib.ScInitFunc(sc.HScName, sc.HFuncInit);
	params: sc.MutableInitParams = new sc.MutableInitParams();
}

export class InitContext {
	params: sc.ImmutableInitParams = new sc.ImmutableInitParams();
	state: sc.MutableIncCounterState = new sc.MutableIncCounterState();
}

export class LocalStateInternalCallCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncLocalStateInternalCall);
}

export class LocalStateInternalCallContext {
	state: sc.MutableIncCounterState = new sc.MutableIncCounterState();
}

export class LocalStatePostCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncLocalStatePost);
}

export class LocalStatePostContext {
	state: sc.MutableIncCounterState = new sc.MutableIncCounterState();
}

export class LocalStateSandboxCallCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncLocalStateSandboxCall);
}

export class LocalStateSandboxCallContext {
	state: sc.MutableIncCounterState = new sc.MutableIncCounterState();
}

export class PostIncrementCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncPostIncrement);
}

export class PostIncrementContext {
	state: sc.MutableIncCounterState = new sc.MutableIncCounterState();
}

export class RepeatManyCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncRepeatMany);
	params: sc.MutableRepeatManyParams = new sc.MutableRepeatManyParams();
}

export class RepeatManyContext {
	params: sc.ImmutableRepeatManyParams = new sc.ImmutableRepeatManyParams();
	state: sc.MutableIncCounterState = new sc.MutableIncCounterState();
}

export class TestVliCodecCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncTestVliCodec);
}

export class TestVliCodecContext {
	state: sc.MutableIncCounterState = new sc.MutableIncCounterState();
}

export class TestVluCodecCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncTestVluCodec);
}

export class TestVluCodecContext {
	state: sc.MutableIncCounterState = new sc.MutableIncCounterState();
}

export class WhenMustIncrementCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncWhenMustIncrement);
	params: sc.MutableWhenMustIncrementParams = new sc.MutableWhenMustIncrementParams();
}

export class WhenMustIncrementContext {
	params: sc.ImmutableWhenMustIncrementParams = new sc.ImmutableWhenMustIncrementParams();
	state: sc.MutableIncCounterState = new sc.MutableIncCounterState();
}

export class GetCounterCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewGetCounter);
	results: sc.ImmutableGetCounterResults = new sc.ImmutableGetCounterResults();
}

export class GetCounterContext {
	results: sc.MutableGetCounterResults = new sc.MutableGetCounterResults();
	state: sc.ImmutableIncCounterState = new sc.ImmutableIncCounterState();
}

export class GetVliCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewGetVli);
	params: sc.MutableGetVliParams = new sc.MutableGetVliParams();
	results: sc.ImmutableGetVliResults = new sc.ImmutableGetVliResults();
}

export class GetVliContext {
	params: sc.ImmutableGetVliParams = new sc.ImmutableGetVliParams();
	results: sc.MutableGetVliResults = new sc.MutableGetVliResults();
	state: sc.ImmutableIncCounterState = new sc.ImmutableIncCounterState();
}

export class GetVluCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewGetVlu);
	params: sc.MutableGetVluParams = new sc.MutableGetVluParams();
	results: sc.ImmutableGetVluResults = new sc.ImmutableGetVluResults();
}

export class GetVluContext {
	params: sc.ImmutableGetVluParams = new sc.ImmutableGetVluParams();
	results: sc.MutableGetVluResults = new sc.MutableGetVluResults();
	state: sc.ImmutableIncCounterState = new sc.ImmutableIncCounterState();
}

export class ScFuncs {
    static callIncrement(ctx: wasmlib.ScFuncCallContext): CallIncrementCall {
        return new CallIncrementCall();
    }

    static callIncrementRecurse5x(ctx: wasmlib.ScFuncCallContext): CallIncrementRecurse5xCall {
        return new CallIncrementRecurse5xCall();
    }

    static endlessLoop(ctx: wasmlib.ScFuncCallContext): EndlessLoopCall {
        return new EndlessLoopCall();
    }

    static increment(ctx: wasmlib.ScFuncCallContext): IncrementCall {
        return new IncrementCall();
    }

    static incrementWithDelay(ctx: wasmlib.ScFuncCallContext): IncrementWithDelayCall {
        let f = new IncrementWithDelayCall();
        f.func.setPtrs(f.params, null);
        return f;
    }

    static init(ctx: wasmlib.ScFuncCallContext): InitCall {
        let f = new InitCall();
        f.func.setPtrs(f.params, null);
        return f;
    }

    static localStateInternalCall(ctx: wasmlib.ScFuncCallContext): LocalStateInternalCallCall {
        return new LocalStateInternalCallCall();
    }

    static localStatePost(ctx: wasmlib.ScFuncCallContext): LocalStatePostCall {
        return new LocalStatePostCall();
    }

    static localStateSandboxCall(ctx: wasmlib.ScFuncCallContext): LocalStateSandboxCallCall {
        return new LocalStateSandboxCallCall();
    }

    static postIncrement(ctx: wasmlib.ScFuncCallContext): PostIncrementCall {
        return new PostIncrementCall();
    }

    static repeatMany(ctx: wasmlib.ScFuncCallContext): RepeatManyCall {
        let f = new RepeatManyCall();
        f.func.setPtrs(f.params, null);
        return f;
    }

    static testVliCodec(ctx: wasmlib.ScFuncCallContext): TestVliCodecCall {
        return new TestVliCodecCall();
    }

    static testVluCodec(ctx: wasmlib.ScFuncCallContext): TestVluCodecCall {
        return new TestVluCodecCall();
    }

    static whenMustIncrement(ctx: wasmlib.ScFuncCallContext): WhenMustIncrementCall {
        let f = new WhenMustIncrementCall();
        f.func.setPtrs(f.params, null);
        return f;
    }

    static getCounter(ctx: wasmlib.ScViewCallContext): GetCounterCall {
        let f = new GetCounterCall();
        f.func.setPtrs(null, f.results);
        return f;
    }

    static getVli(ctx: wasmlib.ScViewCallContext): GetVliCall {
        let f = new GetVliCall();
        f.func.setPtrs(f.params, f.results);
        return f;
    }

    static getVlu(ctx: wasmlib.ScViewCallContext): GetVluCall {
        let f = new GetVluCall();
        f.func.setPtrs(f.params, f.results);
        return f;
    }
}
